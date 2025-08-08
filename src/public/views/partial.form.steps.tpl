{{define "steps"}}
<details>
<summary>{{.EmbedModel.ContentType}}</summary>
<ul>
    {{if eq 1 2}}
    <input type="hidden" name="{{.EmbedModel.ContentType}}_id" value="{{.EmbedModel.ID}}">
    <li><span>screenwriting steps</span><label class="switch"><input type="checkbox" name="{{.EmbedModel.ContentType}}_screenwriting_steps" id="{{.EmbedModel.ContentType}}_screenwriting_steps"><span class="slider round"></span></label></li>
    <li><span>screenwriting get prompt input</span><label class="switch"><input type="checkbox" name="{{.EmbedModel.ContentType}}_screenwriting_get_prompt_input" id="{{.EmbedModel.ContentType}}_screenwriting_get_prompt_input"><span class="slider round"></span></label></li>
    <li><span>screenwriting get prompt output</span><label class="switch"><input type="checkbox" name="{{.EmbedModel.ContentType}}_screenwriting_get_prompt_output" id="{{.EmbedModel.ContentType}}_screenwriting_get_prompt_output"><span class="slider round"></span></label></li>
    <li><span>screenwriting output</span><label class="switch"><input type="checkbox" name="{{.EmbedModel.ContentType}}_screenwriting_output" id="{{.EmbedModel.ContentType}}_screenwriting_output"><span class="slider round"></span></label></li>
    <li><span>container swap</span><label class="switch"><input type="checkbox" name="{{.EmbedModel.ContentType}}_container_swap" id="{{.EmbedModel.ContentType}}_container_swap"><span class="slider round"></span></label></li>
    <li><span>generate audio</span><label class="switch"><input type="checkbox" name="{{.EmbedModel.ContentType}}_generate_audio" id="{{.EmbedModel.ContentType}}_generate_audio"><span class="slider round"></span></label></li>
    <li><span>generate lipsync</span><label class="switch"><input type="checkbox" name="{{.EmbedModel.ContentType}}_generate_lipsync" id="{{.EmbedModel.ContentType}}_generate_lipsync"><span class="slider round"></span></label></li>
    <li><span>generate thumbnails</span><label class="switch"><input type="checkbox" name="{{.EmbedModel.ContentType}}_generate_thumbnails" id="{{.EmbedModel.ContentType}}_generate_thumbnails"><span class="slider round"></span></label></li>
    <li><span>generate background context images</span><label class="switch"><input type="checkbox" name="{{.EmbedModel.ContentType}}_generate_background_context" id="{{.EmbedModel.ContentType}}_generate_background_context"><span class="slider round"></span></label></li>
    <li><span>generate background</span><label class="switch"><input type="checkbox" name="{{.EmbedModel.ContentType}}_generate_background" id="{{.EmbedModel.ContentType}}_generate_background"><span class="slider round"></span></label></li>
    <li><span>ffmpeg lipsync post</span><label class="switch"><input type="checkbox" name="{{.EmbedModel.ContentType}}_ffmpeg_lipsync_post" id="{{.EmbedModel.ContentType}}_ffmpeg_lipsync_post"><span class="slider round"></span></label></li>
    <li><span>ffmpeg merge</span><label class="switch"><input type="checkbox" name="{{.EmbedModel.ContentType}}_ffmpeg_merge" id="{{.EmbedModel.ContentType}}_ffmpeg_merge"><span class="slider round"></span></label></li>
    <li><span>publish video</span><label class="switch"><input type="checkbox" name="{{.EmbedModel.ContentType}}_publish_video" id="{{.EmbedModel.ContentType}}_publish_video"><span class="slider round"></span></label></li>
    <li><span>publish thumbnails</span><label class="switch"><input type="checkbox" name="{{.EmbedModel.ContentType}}_publish_thumbnail" id="{{.EmbedModel.ContentType}}_publish_thumbnail"><span class="slider round"></span></label></li>
    <li><span>publish metadata</span><label class="switch"><input type="checkbox" name="{{.EmbedModel.ContentType}}_publish_metadata" id="{{.EmbedModel.ContentType}}_publish_metadata"><span class="slider round"></span></label></li>
    {{else}}
    <input type="hidden" name="{{.EmbedModel.ContentType}}_id" value="{{.EmbedModel.ID}}">
    <li>{{template "element.toggle" .ScreenwritingStart}}</li>
    <li>{{template "element.toggle" .ScreenwritingGetPromptInput}}</li>
    <li>{{template "element.toggle" .ScreenwritingGetPromptOutput}}</li>
    <li>{{template "element.toggle" .ScreenwritingOutput}}</li>
    <li>{{template "element.toggle" .ContainerSwap}}</li>
    <li>{{template "element.toggle" .GenerateAudio}}</li>
    <li>{{template "element.toggle" .GenerateLipsync}}</li>
    <li>{{template "element.toggle" .GenerateThumbnails}}</li>
    <li>{{template "element.toggle" .GenerateBackgroundCountext}}</li>
    <li>{{template "element.toggle" .GenerateBackground}}</li>
    <li>{{template "element.toggle" .FFMPEGLipsyncPost}}</li>
    <li>{{template "element.toggle" .FFMPEGMerge}}</li>
    <li>{{template "element.toggle" .PublishVideo}}</li>
    <li>{{template "element.toggle" .PublishThumbnail}}</li>
    <li>{{template "element.toggle" .PublishMetadata}}</li>
    {{end}}
</ul>
</details>
{{end}}