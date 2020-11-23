# ravlik-single
Project as POC for AI of how to make correct decisions without knowing the dependency of factors.

There is a couple of colonies of Ravliks (Ukrainian name of Snail)
Ravlik can have a next states: NEW(Just born), IDLE(nothing to do), PARENTING(new ravliks making), DEFENCE(protect from aggressive area), DEAD

Apart from Ravlik there are aggressive Area with next states(SUNNY, CLOUDY, STORMY) There are probabilities how area states come after each other(Ravlik don't know about that). STORMY kills a lot of ravliks and it usually comes after CLOUDY state.

Process like Turn based strategy: Area makes a turn, then Ravlik colonies make their turn. 

Goal of POC is to understand how Ravlik colony survives without IF/ELSE statements. 

As a result: 4 colonies, each has own algorithm 
- RandomRavlik, RandomRavlik2: Use random state change with some probability 
- RavlikWithMemory: Saves decisions, decides randomly at first attempt, saves experience. Then use experience under the dame circumstances.
- RavlikShareWithKids: Share experience with Kids, they will know more.

All colonies have 1000000 ravliks at start, and during 100 turns try to survive.

How to run: 
```go run . ```

Results [here](https://github.com/kosyuk/ravlik-single/blob/master/output.txt)
