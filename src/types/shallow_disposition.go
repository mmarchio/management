package types

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	merrors "github.com/mmarchio/management/errors"
	"github.com/mmarchio/management/models"
)

func NewShallowDisposition(id *string) ShallowDisposition {
	c := ShallowDisposition{}
	c.New(id)
	c.ShallowModel.ContentType = "shallowdisposition"
	c, _ = ValidateShallowDisposition(c)
	return c
} 

func NewShallowDispositionModelContent() models.ShallowContent {
	c := models.ShallowContent{}
	c.ShallowModel.ContentType = "disposition"
	return c
}

func NewShallowDispositionTypeContent() ShallowContent {
	c := ShallowContent{}
	c.ShallowModel.ContentType = "disposition"
	return c
}

type ShallowDisposition struct {
	ShallowModel
	ID 						DispositionID 	`json:"id"`
	Name 					string 			`form:"name" json:"name"`
	MinDuration 			int64 			`form:"min_duration" json:"min_duration"`
	MaxDuration 			int64 			`form:"max_duration" json:"max_duration"`
	AdvertisementDuration 	int64 			`form:"advertisement_duration" json:"advertisement_duration"`
	EntitlementsModel 		string 			`form:"entitlements" json:"entitlements_model"`
	VerificationModel 		string 			`form:"verification" json:"verification_model"`
	BypassModel 			string 			`form:"bypass" json:"bypass_model"`
}

func (c ShallowDisposition) ToContent() (*Content, error) {
	m := Content{}
	m.Model = m.Model.FromShallowModel(c.ShallowModel)
	b, err := json.Marshal(c)
	if err != nil {
		return nil, merrors.JSONMarshallingError{}.Wrap(err)
	}
	m.Content = string(b)
	return &m, nil
}

func (c ShallowDisposition) Expand(ctx context.Context) (*Disposition, error) {
	r := Disposition{}
	r.Model.ID = c.ShallowModel.ID
	r.Model.CreatedAt = c.ShallowModel.CreatedAt
	r.Model.UpdatedAt = c.ShallowModel.UpdatedAt
	r.Model.ContentType = c.ShallowModel.ContentType
	r.ID = DispositionID(c.ID)
	r.Name = c.Name
	r.MinDuration = c.MinDuration
	r.MaxDuration = c.MaxDuration
	r.AdvertisementDuration = c.AdvertisementDuration
	entitlements := ShallowEntitlements{}
	entitlements.ShallowModel.ID = c.EntitlementsModel
	entitlementsModel, err := entitlements.Expand(ctx);
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(err)
	}
	r.EntitlementsModel = *entitlementsModel
	steps := ShallowSteps{}
	steps.ShallowModel.ID = c.VerificationModel
	steps.ID = StepsID(c.VerificationModel)
	steps.ShallowModel.ContentType = "shallowverification"
	verificationModel, err := steps.Expand(ctx)
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(err)
	}
	r.VerificationModel = *verificationModel
	steps.ShallowModel.ID = c.BypassModel
	steps.ID = StepsID(c.BypassModel)
	steps.ContentType = "shallowbypass"
	bypassModel, err := steps.Expand(ctx)
	if err != nil {
		return nil, merrors.ContentGetError{}.Wrap(err)
	}
	r.BypassModel = *bypassModel
	return &r, nil
}

func (c *ShallowDisposition) New(id *string) {
	if id != nil {
		c.ShallowModel.ID = *id
	} else {
		c.ShallowModel.ID = uuid.NewString()
	}
	c.ID = DispositionID(c.ShallowModel.ID)
	c.ShallowModel.CreatedAt = time.Now()
	c.ShallowModel.UpdatedAt = c.ShallowModel.CreatedAt
	c.BypassModel = ""
	c.VerificationModel = ""
}

func (c ShallowDisposition) List(ctx context.Context) ([]ShallowDisposition, error) {
	content := NewDispositionModelContent()
	contents, err := content.List(ctx)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.ShallowModel.ContentType}.Wrap(err)
	}
	cuts := make([]ShallowDisposition, 0)
	for _, model := range contents {
		cut := NewShallowDisposition(nil)
		cut, err = cut.Unmarshal(model.Content)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "Disposition", Function: "List"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c ShallowDisposition) ListBy(ctx context.Context, key string, value interface{}) ([]ShallowDisposition, error) {
	content := NewShallowDispositionModelContent()
	contents, err := content.ListBy(ctx, key, value)
	if err != nil {
		return nil, merrors.ContentListError{Info: c.ShallowModel.ContentType}.Wrap(err)
	}
	cuts := make([]ShallowDisposition, 0)
	for _, model := range contents {
		cut := ShallowDisposition{}
		err = json.Unmarshal([]byte(model.Content), &cut)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "ShallowDisposition", Function: "List"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c *ShallowDisposition) Get(ctx context.Context) error {
	content := NewShallowDispositionTypeContent()
	content.ShallowModel.ContentType = "disposition"
	content.ShallowModel.ID = c.ShallowModel.ID
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.ContentGetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "Disposition", Function: "Get"}.Wrap(err)
	}
	return nil
}

func (c ShallowDisposition) Set(ctx context.Context) error {
	content := NewShallowDispositionTypeContent()
	content.FromType(c)
	content.ShallowModel.ID = c.ShallowModel.ID
	err := content.Set(ctx)
	if err != nil {
		return merrors.ContentSetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	return nil
}

func (c ShallowDisposition) Delete(ctx context.Context) error {
	content := NewShallowDispositionTypeContent()
	content.FromType(c)
	content.ShallowModel.ID = c.ShallowModel.ID
	if err := content.Delete(ctx); err != nil {
		return merrors.ContentDeleteError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	return nil
}

func (c ShallowDisposition) GetID() string {
	return c.ShallowModel.ID
}

func (c ShallowDisposition) GetContentType() string {
	return c.ShallowModel.ContentType
}

func (c ShallowDisposition) SetID() (ShallowDisposition, error) {
	var err error
	c.ID = DispositionID(c.ShallowModel.ID)
	if err != nil {
		return c, merrors.IDSetError{Info: "shallow_disposition"}.Wrap(err)
	}
	return c, nil
}

func (c ShallowDisposition) Unmarshal(j string) (ShallowDisposition, error) {
	model := models.ShallowDisposition{}
	model.ShallowModel.ContentType = "disposition"
	if err := json.Unmarshal([]byte(j), &model); err != nil {
		return c, merrors.JSONUnmarshallingError{Info: j, Package: "types", Struct: "Disposition", Function: "Unmarshal"}.Wrap(err)
	}
	c.ShallowModel.FromModel(model.ShallowModel)

	d, err := c.SetID()
	if err != nil {
		return c, merrors.IDSetError{Info: "disposition"}.Wrap(err)
	}
	c = d
	c.Name = model.Name
	c.MinDuration = model.MinDuration
	c.MaxDuration = model.MaxDuration
	c.AdvertisementDuration = model.AdvertisementDuration
	c.EntitlementsModel = model.EntitlementsModel
	c.VerificationModel = model.VerificationModel
	c.BypassModel = model.BypassModel
	return c, nil
}

func (c ShallowDisposition) Bind(e echo.Context) (ShallowDisposition, error) {
	var err error
	c.BypassModel = e.FormValue("bypass")
	c.VerificationModel = e.FormValue("verification")
	c.EntitlementsModel = e.FormValue("entitlements")
	return c, err
}

func ValidateShallowDisposition(d ShallowDisposition) (ShallowDisposition, error) {
	if d.EntitlementsModel != "" {
		_, err := uuid.Parse(d.EntitlementsModel)
		if err != nil {
			return d, err
		}
	}
	if d.BypassModel != "" {
		_, err := uuid.Parse(d.BypassModel)
		if err != nil {
			return d, err
		}
	}
	if d.VerificationModel != "" {
		_, err := uuid.Parse(d.VerificationModel)
		if err != nil {
			return d, err
		}
	}
	return d, nil
}

func (c ShallowDisposition) IsShallowModel() bool {
	return true
}
