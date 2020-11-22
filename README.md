# ravlik-single
Project as POC for AI of how to make correct desicions without knowing the dependency of factors.

There are a couple of colonies of Ravliks (Ukrainian name of Snail)
Ravlik can have a next states: NEW(Just born), IDLE(nothing to do), PARENTING(new ravliks making), DEFENCE(protect from agressive area), DEAD

Apart from Ravlik there are aggressive Area with next states(SUNNY, CLOUDY, STORMY) There are a probabilities how area states come from each other(Ravlik don't know about that). STORMY kills a lot of ravliks and it usually comes after CLOUDY state.

Process lite Turn based strategy: Area makes a turn, then ravlik colonies make their turn. 

Goal if POC is understand how survive Ravlik colony without IF/ELSE statements. 

As a result: 4 colonies, each has own algorythm 
- RandomRavlik, RandomRavlik2: Use random state change with some probability 
- RavlikWithMemory: Uses random with first attempt, save experience, then when same circumstances use experience, then random
- RavlikShareWithKids: Share experience with Kids, they will know more.

