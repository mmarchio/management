package types

type GenerateAudioTemplate struct {
	Generator 	KokoroGenerator `json:"3"`
	Speaker		KokoroSpeaker 	`json:"4"`
	Save		KokoroSaveAudio `json:"5"`
}

type KokoroGenerator struct {
	Inputs 		KokoroGeneratorInputs	`json:"inputs"`
	ClassType 	string					`json:"class_type"`
	Meta 		KokoroMeta				`json:"_meta"`
}

type KokoroSpeaker struct {
	Inputs KokoroSpeakerInputs `json:"inputs"`
	ClassType string `json:"class_type"`
	Meta KokoroMeta `json:"_meta"`	
}

type KokoroSaveAudio struct {
	Inputs KokoroSaveAudioInputs `json:"inputs"`
	ClassType string `json:"class_type"`
	Meta KokoroMeta `json:"_meta"`
}

type KokoroSaveAudioInputs struct {}

type KokoroGeneratorInputs struct {}

type KokoroSpeakerInputs struct {}

type KokoroMeta struct {}


