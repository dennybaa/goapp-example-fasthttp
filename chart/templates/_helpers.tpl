{{/* vim: set filetype=mustache: */}}

{{/*
  Library mode of dysnix/app has not been tested yet,
  since the direct mode is mostly used so far.

  Thus we need a few workarounds unless I make a PR :)
*/}}

{{/* Fix mode detection */}}
{{- define "app.chart.mode" -}}
  {{- ternary "direct" "library" (and (empty .Subcharts.app) (eq .Chart.IsRoot true)) -}}
{{- end -}}

{{/*
  Function renders resource/resources template such as deployment or configmaps
  into a specific file.

  For example if we want to render deployment resource we need to place the bellow
  code into deployment.yaml:

  ```
  {{- include "app.template" . -}}
  ```

  Placing the include statement into the specific file (such as deployment.yaml)
  is essential since it:
    - Defines which resource to include "deployment"
    - Attaches the resource rendering to this specific file, which results in the correct
      source line to Helm (# Source: demo/templates/deployment.yaml). Which is essential
      for debugging.

*/}}
{{- define "app.template" -}}
  {{- $resource := .Template.Name | base | trimSuffix ".yaml" -}}
  {{- $conventionBroken := dict "tls-secret" "ingress.tls-secret" -}}

  {{- range $_, $component := concat (list "") $.Values.app.components -}}
    {{- $values := ternary $.Values (get $.Values "component") (eq $component "") | default dict -}}
    {{- $r := get $conventionBroken $resource | default $resource -}}
    {{- include (printf "app.%s" $r) (dict "component" $component "values" $values "top" $) -}}
  {{- end -}}
{{- end -}}
