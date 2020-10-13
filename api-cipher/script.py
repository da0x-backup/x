import sys

def main():
    print("Hello from Python!")
    print('Number of arguments:', len(sys.argv), 'arguments.')
    print('Argument List:', str(sys.argv))
    if len(sys.argv) < 2:
        raise ValueError("Not enough arguments")

if __name__ == "__main__":
    main()

