data: list[str] = []

with open('./structs.go', 'r') as s:
    data = [i for i in s if not i.strip().startswith('//')]

with open('./s.go', 'w') as w: w.writelines(data)
