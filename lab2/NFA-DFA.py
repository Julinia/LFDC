import pandas as pd

nonterminal_symbol = ['q0', 'q1', 'q2', 'q3', 'q4']
terminal_symbol = ['a', 'b']

nfa = {
    'S': {'a': 'A'},
    'A': {'b': 'AB'},
    'B': {'a': 'C',
          'b': 'S'},
    'C': {'a': 'D'},
    'D': {'a': 'S'},
}

print(pd.DataFrame.from_dict(nfa).transpose())

def getStates(nfa):
	states = []
	#'nter' stands for the non-terminal symbol; 'ter' for terminal ones
	for nter in nfa:
		states.append(nter)
	for nter in nfa:
		for ter in nfa[nter]:
			if not nfa[nter][ter] in states:
				states.append(nfa[nter][ter])
	return states

states = getStates(nfa)

def convertNFAtoDFA(nfa, states, terminal_symbol):
	for new_states in states:
		if not new_states in nfa:
			nfa[new_states] = {}
			new_ones = list(new_states)
			for symbol in terminal_symbol:
				terminal = []
				for path in new_ones:
					if symbol in nfa[path]:
						terminal.append(nfa[path][symbol])
				nfa[new_states][symbol] = ''.join(set(''.join(terminal)))
				states.append(''.join(set(''.join(terminal))))
	return nfa

dfa = convertNFAtoDFA(nfa, states, terminal_symbol)

DataFrame = pd.DataFrame(dfa)
print(DataFrame.transpose())