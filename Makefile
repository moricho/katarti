goimports:
	find -E . -type f -iregex ".*\.(go)" | xargs -I _ goimports -local "github.com/moricho" -w _
