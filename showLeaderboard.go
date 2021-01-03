{{/*---script 2 generate the leaderboard---*/}}
{{if ($lb:= sdict (or (dbGet 0 "LeaderBoard").Value sdict))}}{{/*load leaderboard data if availble*/}}
{{range $i, $e:= $lb.entries}}{{/*show leaderboard*/}}
{{(add $i 1)}}. **{{.username}}** LVL {{.lvl}}{{end}}
{{else}}
no entries yet
{{end}}
{{/*---end of script 2---*/}}