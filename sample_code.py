import numpy as np


def print_transpose(array):
    """Return the transposed array."""
    array = np.array(array)
    print(array.T)


if __name__ == '__main__':
    print_transpose([[1,2,3], [4,5,6]])
