Naive golang port of [Questlib](https://github.com/Rothens/QuestLib) by Rothens 
```
➜  questlib-go git:(master) ✗ go run proof.go
QuestDef{id=0, description='Write your quest details here', ongoing='Hey, this is easy! Don't mess around!', onfinished='Yaay, you finished the quest!'}
QuestRequest{subject=0, action='kill', amount='10'}
QuestRequest{subject=0, action='kill', amount='10'}
QuestDef{id=1, description='Write your quest details here', ongoing='Hey, This is not yet finished!', onfinished='Yaay, done'}
QuestRequest{subject=0, action='gather', amount='10'}
```