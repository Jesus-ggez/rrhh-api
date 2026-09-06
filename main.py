data: list[str] = []

target: str = 'h2'
with open(f'./{target}.stub.go', 'r') as s:
    data = [i for i in s if not i.strip().startswith('//')]

with open(f'./{target}.h.go', 'w') as w: w.writelines(data)
