import json
# person = '{"name":"Eshaan","age":19}'
# person_dict = json.loads(person)
# print(person_dict)
# with open('data.json') as f:
#     data = json.load(f)
# print(data)
person_json = {"name":"Eshaan","age":19}
# person_json_f = json.dumps(person_json)
# print(person_json_f)
# with open('person.json','w') as json_file:
#     json.dump(person_json, json_file)
with open('data.json') as f:
    data = json.load(f)
print(json.dumps(data, indent=4,sort_keys=True)) # pretty print format