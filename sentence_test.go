package conllu

import (
	"strings"
	"testing"
)

func TestSentenceString(t *testing.T) {

	input := `# sent_id = weblog-juancole.com_juancole_20051126063000_ENG_20051126_063000-0009
# text = Although the announcement was probably made to show progress in identifying and breaking up terror cells, I don't find the news that the Baathists continue to penetrate the Iraqi government very hopeful.
1	Although	although	SCONJ	IN	_	6	mark	6:mark	_
2	the	the	DET	DT	Definite=Def|PronType=Art	3	det	3:det	_
3	announcement	announcement	NOUN	NN	Number=Sing	6	nsubj:pass	6:nsubj:pass|8:nsubj:xsubj	_
4	was	be	AUX	VBD	Mood=Ind|Number=Sing|Person=3|Tense=Past|VerbForm=Fin	6	aux:pass	6:aux:pass	_
5	probably	probably	ADV	RB	_	6	advmod	6:advmod	_
6	made	make	VERB	VBN	Tense=Past|VerbForm=Part|Voice=Pass	21	advcl	21:advcl:although	_
7	to	to	PART	TO	_	8	mark	8:mark	_
8	show	show	VERB	VB	VerbForm=Inf	6	xcomp	6:xcomp	_
9	progress	progress	NOUN	NN	Number=Sing	8	obj	8:obj	_
10	in	in	SCONJ	IN	_	11	mark	11:mark	_
11	identifying	identify	VERB	VBG	VerbForm=Ger	9	acl	9:acl:in	_
12	and	and	CCONJ	CC	_	13	cc	13:cc	_
13	breaking	break	VERB	VBG	VerbForm=Ger	11	conj	9:acl:in|11:conj:and	_
14	up	up	ADP	RP	_	13	compound:prt	13:compound:prt	_
15	terror	terror	NOUN	NN	Number=Sing	16	compound	16:compound	_
16	cells	cell	NOUN	NNS	Number=Plur	11	obj	11:obj|13:obj	SpaceAfter=No
17	,	,	PUNCT	,	_	21	punct	21:punct	_
18	I	I	PRON	PRP	Case=Nom|Number=Sing|Person=1|PronType=Prs	21	nsubj	21:nsubj	_
19-20	don't	_	_	_	_	_	_	_	_
19	do	do	AUX	VBP	Mood=Ind|Tense=Pres|VerbForm=Fin	21	aux	21:aux	_
20	n't	not	PART	RB	_	21	advmod	21:advmod	_
21	find	find	VERB	VB	VerbForm=Inf	0	root	0:root	_
22	the	the	DET	DT	Definite=Def|PronType=Art	23	det	23:det	_
23	news	news	NOUN	NN	Number=Sing	21	obj	21:obj|34:nsubj:xsubj	_
24	that	that	SCONJ	IN	_	27	mark	27:mark	_
25	the	the	DET	DT	Definite=Def|PronType=Art	26	det	26:det	_
26	Baathists	Baathists	PROPN	NNPS	Number=Plur	27	nsubj	27:nsubj|29:nsubj:xsubj	_
27	continue	continue	VERB	VBP	Mood=Ind|Tense=Pres|VerbForm=Fin	23	acl	23:acl:that	_
28	to	to	PART	TO	_	29	mark	29:mark	_
29	penetrate	penetrate	VERB	VB	VerbForm=Inf	27	xcomp	27:xcomp	_
30	the	the	DET	DT	Definite=Def|PronType=Art	32	det	32:det	_
31	Iraqi	iraqi	ADJ	JJ	Degree=Pos	32	amod	32:amod	_
32	government	government	NOUN	NN	Number=Sing	29	obj	29:obj	_
33	very	very	ADV	RB	_	34	advmod	34:advmod	_
34	hopeful	hopeful	ADJ	JJ	Degree=Pos	21	xcomp	21:xcomp	SpaceAfter=No
35	.	.	PUNCT	.	_	21	punct	21:punct	_
`

	sentences, err := Parse(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	if len(sentences) != 1 {
		t.Errorf("Expected only 1 sentence, got %d", len(sentences))
	}
	
	want := "Although the announcement was probably made to show progress in identifying and breaking up terror cells, I don't find the news that the Baathists continue to penetrate the Iraqi government very hopeful."
	got := sentences[0].String()

	if got != want {
		t.Errorf("Invalid sentence\n\tWant %q\n\tGot: %q", want, got)
	}
}

