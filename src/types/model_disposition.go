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

func NewDisposition(id *string) Disposition {
	c := Disposition{}
	c.New(id)
	c.Model.ContentType = "disposition"
	c, _ = ValidateDisposition(c)
	return c
} 

func NewShallowDisposition(id *string) ShallowDisposition {
	c := ShallowDisposition{}
	c.New(id)
	c.ShallowModel.ContentType = "disposition"
	c, _ = ValidateShallowDisposition(c)
	return c
} 

func NewDispositionModelContent() models.Content {
	c := models.Content{}
	c.Model.ContentType = "disposition"
	return c
}

func NewDispositionTypeContent() Content {
	c := Content{}
	c.Model.ContentType = "disposition"
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

type Disposition struct {
	Model
	ID 						DispositionID 	`json:"id"`
	Name 					string 			`form:"name" json:"name"`
	MinDuration 			int64 			`form:"min_duration" json:"min_duration"`
	MaxDuration 			int64 			`form:"max_duration" json:"max_duration"`
	AdvertisementDuration 	int64 			`form:"advertisement_duration" json:"advertisement_duration"`
	EntitlementsModel 		Entitlements 	`form:"entitlements" json:"entitlements_model"`
	VerificationModel 		Steps 			`form:"verification" json:"verification_model"`
	BypassModel 			Steps 			`form:"bypass" json:"bypass_model"`
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

func (c *Disposition) New(id *string) {
	if id != nil {
		c.Model.ID = *id
	} else {
		c.Model.ID = uuid.NewString()
	}
	c.ID = DispositionID(c.Model.ID)
	c.Model.CreatedAt = time.Now()
	c.Model.UpdatedAt = c.Model.CreatedAt
	c.BypassModel.EmbedModel.New("bypass")
	c.VerificationModel.EmbedModel.New("verification")
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

func (c Disposition) List(ctx context.Context) ([]Disposition, error) {
	content := NewDispositionModelContent()
	contents, err := content.List(ctx)
	if err != nil {
		return nil, merrors.DispositionListError{Info: c.Model.ContentType}.Wrap(err)
	}
	cuts := make([]Disposition, 0)
	for _, model := range contents {
		cut := NewDisposition(nil)
		cut, err = cut.Unmarshal(model.Content)
		if err != nil {
			return nil, merrors.JSONUnmarshallingError{Info: model.Content, Package: "types", Struct: "Disposition", Function: "List"}.Wrap(err)
		}
		cuts = append(cuts, cut)
	}
	return cuts, nil
}

func (c ShallowDisposition) List(ctx context.Context) ([]ShallowDisposition, error) {
	content := NewDispositionModelContent()
	contents, err := content.List(ctx)
	if err != nil {
		return nil, merrors.DispositionListError{Info: c.ShallowModel.ContentType}.Wrap(err)
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

func (c Disposition) ListBy(ctx context.Context, key string, value interface{}) ([]Disposition, error) {
	content := NewDispositionModelContent()
	contents, err := content.ListBy(ctx, key, value)
	if err != nil {
		return nil, merrors.DispositionListError{Info: c.Model.ContentType}.Wrap(err)
	}
	cuts := make([]Disposition, 0)
	for _, model := range contents {
		cut := Disposition{}
		err = json.Unmarshal([]byte(model.Content), &cut)
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
		return nil, merrors.DispositionListError{Info: c.ShallowModel.ContentType}.Wrap(err)
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

func (c *Disposition) Get(ctx context.Context) error {
	content := NewDispositionTypeContent()
	content.Model.ContentType = "disposition"
	content.Model.ID = c.Model.ID
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.DispositionGetError{Info: c.Model.ID}.Wrap(err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "Disposition", Function: "Get"}.Wrap(err)
	}
	return nil
}

func (c *ShallowDisposition) Get(ctx context.Context) error {
	content := NewShallowDispositionTypeContent()
	content.ShallowModel.ContentType = "disposition"
	content.ShallowModel.ID = c.ShallowModel.ID
	content, err := content.Get(ctx)
	if err != nil {
		return merrors.DispositionGetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	err = json.Unmarshal([]byte(content.Content), c)
	if err != nil {
		return merrors.JSONUnmarshallingError{Info: content.Content, Package: "types", Struct: "Disposition", Function: "Get"}.Wrap(err)
	}
	return nil
}

func (c Disposition) Set(ctx context.Context) error {
	content := NewDispositionTypeContent()
	content.FromType(c)
	content.Model.ID = c.Model.ID
	err := content.Set(ctx)
	if err != nil {
		return merrors.DispositionSetError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

func (c ShallowDisposition) Set(ctx context.Context) error {
	content := NewShallowDispositionTypeContent()
	content.FromType(c)
	content.ShallowModel.ID = c.ShallowModel.ID
	err := content.Set(ctx)
	if err != nil {
		return merrors.DispositionSetError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	return nil
}

func (c Disposition) Delete(ctx context.Context) error {
	content := NewDispositionTypeContent()
	content.FromType(c)
	content.Model.ID = c.Model.ID
	if err := content.Delete(ctx); err != nil {
		return merrors.DispositionDeleteError{Info: c.Model.ID}.Wrap(err)
	}
	return nil
}

func (c ShallowDisposition) Delete(ctx context.Context) error {
	content := NewShallowDispositionTypeContent()
	content.FromType(c)
	content.ShallowModel.ID = c.ShallowModel.ID
	if err := content.Delete(ctx); err != nil {
		return merrors.DispositionDeleteError{Info: c.ShallowModel.ID}.Wrap(err)
	}
	return nil
}

func (c Disposition) GetID() string {
	return c.Model.ID
}

func (c ShallowDisposition) GetID() string {
	return c.ShallowModel.ID
}

func (c Disposition) GetContentType() string {
	return c.Model.ContentType
}

func (c ShallowDisposition) GetContentType() string {
	return c.ShallowModel.ContentType
}

func (c Disposition) SetID() (Disposition, error) {
	var err error
	c.ID = DispositionID(c.Model.ID)
	if err != nil {
		return c, merrors.IDSetError{Info: "disposition"}.Wrap(err)
	}
	return c, nil
}

func (c ShallowDisposition) SetID() (ShallowDisposition, error) {
	var err error
	c.ID = DispositionID(c.ShallowModel.ID)
	if err != nil {
		return c, merrors.IDSetError{Info: "shallow_disposition"}.Wrap(err)
	}
	return c, nil
}

func (c Disposition) Unmarshal(j string) (Disposition, error) {
	model := models.Disposition{}
	model.Model.ContentType = "disposition"
	if err := json.Unmarshal([]byte(j), &model); err != nil {
		return c, merrors.JSONUnmarshallingError{Info: j, Package: "types", Struct: "Disposition", Function: "Unmarshal"}.Wrap(err)
	}
	c.Model.FromModel(model.Model)

	d, err := c.SetID()
	if err != nil {
		return c, merrors.IDSetError{Info: "disposition"}.Wrap(err)
	}
	c = d
	c.Name = model.Name
	c.MinDuration = model.MinDuration
	c.MaxDuration = model.MaxDuration
	c.AdvertisementDuration = model.AdvertisementDuration
	c.EntitlementsModel = c.EntitlementsModel.FromModel(model.EntitlementsModel)
	c.VerificationModel = c.VerificationModel.FromModel(model.VerificationModel)
	c.BypassModel = c.BypassModel.FromModel(model.BypassModel)
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

func (c Disposition) Bind(e echo.Context) (Disposition, error) {
	var err error
	c.BypassModel, err = c.BypassModel.Bind(e)
	c.VerificationModel, err = c.VerificationModel.Bind(e)
	c.EntitlementsModel, err = c.EntitlementsModel.Bind(e)
	return c, err
}

func (c ShallowDisposition) Bind(e echo.Context) (ShallowDisposition, error) {
	var err error
	c.BypassModel = e.FormValue("bypass")
	c.VerificationModel = e.FormValue("verification")
	c.EntitlementsModel = e.FormValue("entitlements")
	return c, err
}

func ValidateDisposition(d Disposition) (Disposition, error) {
	d.EntitlementsModel, _ = ValidateEntitlements(d.EntitlementsModel, "entitlements_")
	d.EntitlementsModel.EmbedModel.ContentType = "entitlements"
	d.BypassModel, _ = ValidateSteps(d.BypassModel, "bypass_")
	d.BypassModel.EmbedModel.ContentType = "bypass"
	d.VerificationModel, _ = ValidateSteps(d.VerificationModel, "verification_")
	d.VerificationModel.EmbedModel.ContentType = "verification"
	return d, nil
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