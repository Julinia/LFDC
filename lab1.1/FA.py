import os
import sys
os.environ["PATH"] += os.pathsep + 'C:/Users/Iuliana/Desktop/sem 4/graphviz-2.38/release/bin'
from graphviz import Digraph

f = Digraph('finite_state_machine', filename='F-a.gv')
f.attr(rankdir = 'LR', size = '8,5')

f.attr('node', shape = 'doublecircle')
f.node('G')

f.attr('node', shape = 'circle')
f.edge('S', 'D', label = 'a')
f.edge('D', 'E', label = 'b')
f.edge('E', 'F', label = 'c')
f.edge('F', 'D', label = 'd')
f.edge('E', 'L', label = 'd')
f.edge('L', 'L', label = 'a')
f.edge('L', 'G', label = 'c')

f.attr('node', color = 'white')
f.edge('', 'S')

f.view()

fa = {
    'S': {'a': 'D'},
    'D': {'b': 'E'},
    'E': {'c': 'F',
          'd': 'L'},
    'F': {'d': 'D'},
    'L': {'a': 'L',
          'b': 'L',
          'c': 'G'},
}

def accepts(fa, initial, final, string):
	state = initial
	for i in string:
		if state in fa and i in fa[state]:
			state = fa[state][i]
		else:
			return False

	return state in final

print("Enter the word which should be verified in our grammar:")
print(fa)
while True:
	uInput = input('> ')

	if uInput == '':
		break

	print(accepts(fa, 'S', {'G'}, uInput)) 

