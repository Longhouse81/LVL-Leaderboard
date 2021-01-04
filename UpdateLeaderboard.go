{{/*---script 1 to increase the LVL of a user, and if increased check if the user is able to get a spot in the leaderboard---*/}}
{{/*if condition}}{{/*conditions to increase lvl start here*/}}
{{$amount:= 1}}{{/*the amount we increase the lvl with*/}}
{{$maxEntry:= 10}}{{/*we set the max amount of leaderboard entries*/}}
{{$member:= (getMember .User.ID)}}
{{$uLVL:= (toInt (dbIncr $member.User.ID "uLVL" $amount))}}{{/*increase the LVL of the user with the amount you want*/}}

{{/*update leaderboard if needed*/}}
{{$lb:= ""}}
{{$uEntry:= sdict "uID" $member.User.ID "uString" $member.User.String "uLVL" $uLVL}}
{{if ($lb =sdict (or (dbGet 0 "LeaderBoard").Value sdict))}}{{/*load leaderboard data if availble*/}}
	{{$newLB:= cslice}}
	{{$lowest:= 0}}
	{{$isSet:= 0}}
	{{if gt $uLVL $lb.lowest}}
		{{range $i, $e:= $lb.entries}}
			{{- if and (not (eq .uID $member.User.ID)) (lt (len $newLB) $maxEntry) $isSet}}
				{{- $newLB = ($newLB.Append .)}}
				{{- $lowest = .uLVL}}
			{{- else if gt $uLVL .uLVL}}
				{{- if or (eq .uID $member.User.ID) (ge (len $newLB) $maxEntry)}}
					{{- $newLB = ($newLB.Append $uEntry)}}
					{{- $lowest = $uLVL}}
				{{- else if lt (len $newLB) $maxEntry}}
					{{- $newLB = ($newLB.AppendSlice (cslice $uEntry .))}}
					{{- $lowest = .uLVL}}
				{{- end}}
				{{- $isSet = 1}}
			{{- end -}}
		{{end}}
	{{else if lt $lb.length $maxEntry}}
		{{$newLB = ($lb.entries).Append $uEntry}}
		{{$lowest = $uLVL}}
		{{$isSet = 1}}
	{{end}}
	{{if $isSet}}
		{{$lb = sdict "length" (len $newLB) "lowest" $lowest "entries" $newLB}}
	{{end}}
{{else}}
	{{$lb = sdict "length" 1 "lowest" $uLVL "entries" (cslice $uEntry)}}
{{end}}
{{if $lb}}
	{{dbSet 0 "LeaderBoard" $lb}}
{{end}}
{{/*end}}{{/*conditions to increase lvl end here*/}}
{{/*---end of script 1---*/}}