x,y=3,18
step = 3
for i in range(10):
    if i%2==0:
        y+=step
    else:
        x+=step
    print(f'{x},{y}')


for i in range(10):
    if i%2==0:
        y-=step
    else:
        x+=step
    print(f'{x},{y}')


for i in range(10):
    if i%2==0:
        y-=step
    else:
        x-=step
    print(f'{x},{y}')

for i in range(10):
    if i%2==0:
        y+=step
    else:
        x-=step
    print(f'{x},{y}')
