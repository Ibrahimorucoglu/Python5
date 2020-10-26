male(hummet).
male(qurban).
male(orucali).
male(ibrahim).
male(vusal).

female(xaver).
female(suad).
female(shafa).
female(shehla).

spouse(hummet, xaver).
spouse(qurban, suad).
spouse(orucali, shafa).

father(hummet, orucali).
father(qurban, shafa).
father(qurban, shehla).
father(orucali, ibrahim).
father(orucali, vusal).

mother(xaver, orucali).
mother(suad, shafa).
mother(suad, shehla).
mother(shafa, ibrahim).
mother(shafa, vusal).

brother(A,B) :- male(A), not(A=B), father(F,A), father(F,B), mother(M,A), mother(M,B).
brother(vusal, ibrahim).

sister(A,B) :- female(A), not(A=B), father(F,A), father(F,B), mother(M,A), mother(M,B).
sister(shafa, shehla).

sibling(A,B) :- brother(A,B);sister(A,B).
sibling(shehla, shafa).

parent(A,B) :- father(A,B).
parent(A,B) :- mother(A,B).
parent(orucali, ibrahim).

grandfather(A,B) :- male(A), father(A,F), parent(F,B).
grandfather(qurban, vusal).

grandmother(A,B) :- female(A), mother(A,F), parent(F,B).
grandmother(xaver, ibrahim).
