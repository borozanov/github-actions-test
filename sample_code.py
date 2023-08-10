import numpy as np

def print_transpose(a):
    a = np.array(a)
    print(a.T)


if __name__ == '__main__':
    print_transpose([[1,2,3], [4,5,6]])
