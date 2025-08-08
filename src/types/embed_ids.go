package types

import "github.com/google/uuid"

type JobStatusID string

func (c JobStatusID) New() JobStatusID {
	return JobStatusID(uuid.NewString())
}

func (c JobStatusID) IsNil() bool {
	return string(c) == ""
}

func (c JobStatusID) String() string {
	return string(c)
}

func (c JobStatusID) Scan(s interface{}) (JobStatusID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return JobStatusID(n.String()), nil
}

type JobID string

func (c JobID) New() JobID {
	return JobID(uuid.NewString())
}

func (c JobID) IsNil() bool {
	return string(c) == ""
}

func (c JobID) String() string {
	return string(c)
}

func (c JobID) Scan(s interface{}) (JobID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return JobID(n.String()), nil
}

type PromptID string

func (c PromptID) New() PromptID {
	return PromptID(uuid.NewString())
}

func (c PromptID) IsNil() bool {
	return string(c) == ""
}

func (c PromptID) String() string {
	return string(c)
}

func (c PromptID) Scan(s interface{}) (PromptID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return PromptID(n.String()), nil
}

type StatsID string

func (c StatsID) New() StatsID {
	return StatsID(uuid.NewString())
}

func (c StatsID) IsNil() bool {
	return string(c) == ""
}

func (c StatsID) String() string {
	return string(c)
}

func (c StatsID) Scan(s interface{}) (StatsID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return StatsID(n.String()), nil
}

type AudioOutputID string

func (c AudioOutputID) New() AudioOutputID {
	return AudioOutputID(uuid.NewString())
}

func (c AudioOutputID) IsNil() bool {
	return string(c) == ""
}

func (c AudioOutputID) String() string {
	return string(c)
}

func (c AudioOutputID) Scan(s interface{}) (AudioOutputID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return AudioOutputID(n.String()), nil
}

type FileID string

func (c FileID) New() FileID {
	return FileID(uuid.NewString())
}

func (c FileID) IsNil() bool {
	return string(c) == ""
}

func (c FileID) String() string {
	return string(c)
}

func (c FileID) Scan(s interface{}) (FileID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return FileID(n.String()), nil
}

type SceneID string

func (c SceneID) New() SceneID {
	return SceneID(uuid.NewString())
}

func (c SceneID) IsNil() bool {
	return string(c) == ""
}

func (c SceneID) String() string {
	return string(c)
}

func (c SceneID) Scan(s interface{}) (SceneID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return SceneID(n.String()), nil
}

type VideoOutputID string

func (c VideoOutputID) New() VideoOutputID {
	return VideoOutputID(uuid.NewString())
}

func (c VideoOutputID) IsNil() bool {
	return string(c) == ""
}

func (c VideoOutputID) String() string {
	return string(c)
}

func (c VideoOutputID) Scan(s interface{}) (VideoOutputID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return VideoOutputID(n.String()), nil
}

type RunID string

func (c RunID) New() RunID {
	return RunID(uuid.NewString())
}

func (c RunID) IsNil() bool {
	return string(c) == ""
}

func (c RunID) String() string {
	return string(c)
}

func (c RunID) Scan(s interface{}) (RunID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return RunID(n.String()), nil
}

type VideoLipsyncOutputID string

func (c VideoLipsyncOutputID) New() VideoLipsyncOutputID {
	return VideoLipsyncOutputID(uuid.NewString())
}

func (c VideoLipsyncOutputID) IsNil() bool {
	return string(c) == ""
}

func (c VideoLipsyncOutputID) String() string {
	return string(c)
}

func (c VideoLipsyncOutputID) Scan(s interface{}) (VideoLipsyncOutputID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return VideoLipsyncOutputID(n.String()), nil
}

type VideoTransparencyOutputID string

func (c VideoTransparencyOutputID) New() VideoTransparencyOutputID {
	return VideoTransparencyOutputID(uuid.NewString())
}

func (c VideoTransparencyOutputID) IsNil() bool {
	return string(c) == ""
}

func (c VideoTransparencyOutputID) String() string {
	return string(c)
}

func (c VideoTransparencyOutputID) Scan(s interface{}) (VideoTransparencyOutputID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return VideoTransparencyOutputID(n.String()), nil
}

type VideoBackgroundOutputID string

func (c VideoBackgroundOutputID) New() VideoBackgroundOutputID {
	return VideoBackgroundOutputID(uuid.NewString())
}

func (c VideoBackgroundOutputID) IsNil() bool {
	return string(c) == ""
}

func (c VideoBackgroundOutputID) String() string {
	return string(c)
}

func (c VideoBackgroundOutputID) Scan(s interface{}) (VideoBackgroundOutputID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return VideoBackgroundOutputID(n.String()), nil
}

type VideoLayerMergeOutputID string

func (c VideoLayerMergeOutputID) New() VideoLayerMergeOutputID {
	return VideoLayerMergeOutputID(uuid.NewString())
}

func (c VideoLayerMergeOutputID) IsNil() bool {
	return string(c) == ""
}

func (c VideoLayerMergeOutputID) String() string {
	return string(c)
}

func (c VideoLayerMergeOutputID) Scan(s interface{}) (VideoLayerMergeOutputID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return VideoLayerMergeOutputID(n.String()), nil
}

type ImageThumbnailOutputID string

func (c ImageThumbnailOutputID) New() ImageThumbnailOutputID {
	return ImageThumbnailOutputID(uuid.NewString())
}

func (c ImageThumbnailOutputID) IsNil() bool {
	return string(c) == ""
}

func (c ImageThumbnailOutputID) String() string {
	return string(c)
}

func (c ImageThumbnailOutputID) Scan(s interface{}) (ImageThumbnailOutputID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return ImageThumbnailOutputID(n.String()), nil
}

type ImageBackgroundContextOutputID string

func (c ImageBackgroundContextOutputID) New() ImageBackgroundContextOutputID {
	return ImageBackgroundContextOutputID(uuid.NewString())
}

func (c ImageBackgroundContextOutputID) IsNil() bool {
	return string(c) == ""
}

func (c ImageBackgroundContextOutputID) String() string {
	return string(c)
}

func (c ImageBackgroundContextOutputID) Scan(s interface{}) (ImageBackgroundContextOutputID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return ImageBackgroundContextOutputID(n.String()), nil
}

type SettingsID string

func (c SettingsID) New() SettingsID {
	return SettingsID(uuid.NewString())
}

func (c SettingsID) IsNil() bool {
	return string(c) == ""
}

func (c SettingsID) String() string {
	return string(c)
}

func (c SettingsID) Scan(s interface{}) (SettingsID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return SettingsID(n.String()), nil
}

type TemplateID string

func (c TemplateID) New() TemplateID {
	return TemplateID(uuid.NewString())
}

func (c TemplateID) IsNil() bool {
	return string(c) == ""
}

func (c TemplateID) String() string {
	return string(c)
}

func (c TemplateID) Scan(s interface{}) (TemplateID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return TemplateID(n.String()), nil
}

type DispositionID string

func (c DispositionID) New() DispositionID {
	return DispositionID(uuid.NewString())
}

func (c DispositionID) IsNil() bool {
	return string(c) == ""
}

func (c DispositionID) String() string {
	return string(c)
}

func (c DispositionID) Scan(s interface{}) (DispositionID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return DispositionID(n.String()), nil
}

type EntitlementID string

func (c EntitlementID) New() EntitlementID {
	return EntitlementID(uuid.NewString())
}

func (c EntitlementID) IsNil() bool {
	return string(c) == ""
}

func (c EntitlementID) String() string {
	return string(c)
}

func (c EntitlementID) Scan(s interface{}) (EntitlementID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return EntitlementID(n.String()), nil
}

type StepsID string

func (c StepsID) New() StepsID {
	return StepsID(uuid.NewString())
}

func (c StepsID) IsNil() bool {
	return string(c) == ""
}

func (c StepsID) String() string {
	return string(c)
}

func (c StepsID) Scan(s interface{}) (StepsID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return StepsID(n.String()), nil
}

type ComfyUITemplateID string

func (c ComfyUITemplateID) New() ComfyUITemplateID {
	return ComfyUITemplateID(uuid.NewString())
}

func (c ComfyUITemplateID) IsNil() bool {
	return string(c) == ""
}

func (c ComfyUITemplateID) String() string {
	return string(c)
}

func (c ComfyUITemplateID) Scan(s interface{}) (ComfyUITemplateID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return ComfyUITemplateID(n.String()), nil
}

type ContentID string

func (c ContentID) New() ContentID {
	return ContentID(uuid.NewString())
}

func (c ContentID) IsNil() bool {
	return string(c) == ""
}

func (c ContentID) String() string {
	return string(c)
}

func (c ContentID) Scan(s interface{}) (ContentID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return ContentID(n.String()), nil
}

type SystemPromptID string 

func (c SystemPromptID) New() SystemPromptID {
	return SystemPromptID(uuid.NewString())
}

func (c SystemPromptID) IsNil() bool {
	return string(c) == ""
}

func (c SystemPromptID) String() string {
	return string(c)
}

func (c SystemPromptID) Scan(s interface{}) (SystemPromptID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return SystemPromptID(n.String()), nil
}

type AutomationWorkflowID string

func (c AutomationWorkflowID) New() AutomationWorkflowID {
	return AutomationWorkflowID(uuid.NewString())
}

func (c AutomationWorkflowID) IsNil() bool {
	return string(c) == ""
}

func (c AutomationWorkflowID) String() string {
	return string(c)
}

func (c AutomationWorkflowID) Scan(s interface{}) (AutomationWorkflowID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return AutomationWorkflowID(n.String()), nil
}

type AutomationStepID string

func (c AutomationStepID) New() AutomationStepID {
	return AutomationStepID(uuid.NewString())
}

func (c AutomationStepID) IsNil() bool {
	return string(c) == ""
}

func (c AutomationStepID) String() string {
	return string(c)
}

func (c AutomationStepID) Scan(s interface{}) (AutomationStepID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return AutomationStepID(n.String()), nil
}

type NodeID string

func (c NodeID) New() NodeID {
	return NodeID(uuid.NewString())
}

func (c NodeID) IsNil() bool {
	return string(c) == ""
}

func (c NodeID) String() string {
	return string(c)
}

func (c NodeID) Scan(s interface{}) (NodeID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return NodeID(n.String()), nil
}

type WorkflowID string

func (c WorkflowID) New(id *string) WorkflowID {
	if id != nil {
		return WorkflowID(*id)
	}
	return WorkflowID(uuid.NewString())
}

func (c WorkflowID) IsNil() bool {
	return string(c) == ""
}

func (c WorkflowID) String() string {
	return string(c)
}

func (c WorkflowID) Scan(s interface{}) (WorkflowID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return WorkflowID(n.String()), nil
}

type PromptTemplateID string

func (c PromptTemplateID) New(id *string) PromptTemplateID {
	if id != nil {
		return PromptTemplateID(*id)
	}
	return PromptTemplateID(uuid.NewString())
}

func (c PromptTemplateID) IsNil() bool {
	return string(c) == ""
}

func (c PromptTemplateID) String() string {
	return string(c)
}

func (c PromptTemplateID) Scan(s interface{}) (PromptTemplateID, error) {
	n := uuid.New()
	err := n.Scan(s)
	if err != nil {
		return c, err
	}
	return PromptTemplateID(n.String()), nil
}


