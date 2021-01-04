{{/*---script 1 to increase the LVL of a user, and if increased check if the user is able to get a spot in the leaderboard---*/}}
{{/*if condition}}{{/*conditions to increase lvl start here*/}}
{{$amount:= 1}}{{/*the amount we increase the lvl with*/}}
{{$maxEntry:= 10}}{{/*we set the max amount of leaderboard entries*/}}
{{$member:= (getMember .User.ID)}}
{{$userLevel:= (toInt (dbIncr $member.User.ID "lvl" $amount))}}{{/*increase the LVL of the user with the amount you want*/}}

{{/*update leaderboard if needed*/}}
{{$lb:= ""}}
{{$uEntry:= sdict "userid" $member.User.ID "username" (or $member.Nick $member.User.Username) "lvl" $userLevel}}
{{if ($lb =sdict (or (dbGet 0 "LeaderBoard").Value sdict))}}{{/*load leaderboard data if availble*/}}
	{{$newLB:= cslice}}
	{{$lowest:= 0}}
	{{if or (gt $userLevel $lb.lowest) (lt $lb.length $maxEntry)}}
		{{$isSet:= 0}}
		{{range $i, $e:= $lb.entries}}
			{{- if and (gt $userLevel .lvl) (not $isSet)}}
				{{- if eq .userid $member.User.ID}}
					{{- $newLB = ($newLB.Append $uEntry)}}
					{{- $isSet = 1}}
				{{- else}}
					{{- if lt $maxEntry (len $newLB)}}
						{{- $newLB = ($newLB.AppendSlice (cslice $uEntry .))}}
					{{- else}}
						{{- $newLB = ($newLB.Append $uEntry)}}
					{{- end}}
					{{- $isSet = 1}}
				{{- end}}
			{{- else if and (not (eq .userid $member.User.ID)) (lt $i $maxEntry)}}
				{{- $newLB = ($newLB.Append .)}}
			{{- end}}
			{{- if or (eq (sub (len $lb.entries) 1) $i) (eq (sub $maxEntry 1) $i)}}
				{{- if and (not $isSet) (le (len $newLB) $maxEntry)}}
					{{- $newLB = ($newLB.Append $uEntry)}}
				{{- end}}
				{{- $lowest = (index $newLB (sub (len $newLB) 1)).lvl}}
			{{- end -}}
		{{else}}
			{{$newLB = ($newLB.Append $uEntry)}}
			{{$lowest = $userLevel}}
		{{end}}
		{{$lb = sdict "length" (len $newLB) "lowest" $lowest "entries" $newLB}}
	{{end}}
{{else}}
	{{$lb = sdict "length" 1 "lowest" $userLevel "entries" (cslice $uEntry)}}
{{end}}
{{if $lb}}
	{{dbSet 0 "LeaderBoard" $lb}}
{{end}}
{{/*end}}{{/*conditions to increase lvl end here*/}}
{{/*---end of script 1---*/}}