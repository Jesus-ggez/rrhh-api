data: list[str] = []

target: str = input('target of creation headers: ')
with open(f'./{target}.stub.go', 'r') as s:
    is_tab: bool = False

    for i in s:
        i: str
        if '.?' in i:
            is_tab = not is_tab

        if i.strip().startswith('//'):
            continue

        if is_tab and i.strip():
            i = i[::-1]
            ix: int = i.find(' ')
            i = i[:ix + 1] + ':' + i[ix + 1:]
            i = i[::-1].rstrip() + ','

        data.append(i)

with open(f'./{target}.h.go', 'w') as w: w.writelines(data)
