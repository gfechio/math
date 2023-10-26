import random
import matplotlib.pyplot as plt

size = 100
count = 0

class Prisioner:
    def __init__(self, size):
        self.number = random.randint(0,size)

class Boxes:
     L = list(range(1, size))
     random.shuffle(L)

def pick_box(count):
    box_list = Boxes.L
    for i in range(0,size):
        prisioner = Prisioner(size)
        for box in box_list:
            if prisioner.number == box:
                count = count + 1
                break
            else:
                continue
    return(count)

def pick_box_no_loop(count):
    box_list = Boxes.L
    for i in range(0,size):
        for box in box_list:
            count = count + 1
    return(count)


def __main__(): 
    top_followers = new_profile.sort_values(by='followers', axis=0, ascending=False)[:100]

    with_loop = pick_box(count)
    print(with_loop)
    without_loop = pick_box_no_loop(count)
    print(without_loop)
    
    print(with_loop*100/without_loop)
