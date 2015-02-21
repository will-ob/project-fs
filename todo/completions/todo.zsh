if [[ ! -o interactive ]]; then
    return
fi

compctl -K _todo todo

_todo() {
  local word words completions
  read -cA words
  word="${words[2]}"

  if [ "${#words}" -eq 2 ]; then
    completions="$(todo commands)"
  else
    completions="$(todo completions "${word}")"
  fi

  reply=("${(ps:\n:)completions}")
}
