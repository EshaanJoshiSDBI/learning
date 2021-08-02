import pickle
mydict = {1:'a',2:'b',3:'c'}
pickle_file = open('pickle.txt','wb')
pickle.dump(mydict,pickle_file)
pickle_file = open('pickle.txt','rb')
newdict = pickle.load(pickle_file)
print(newdict)