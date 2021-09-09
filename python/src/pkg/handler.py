from ddtrace import tracer

@tracer.wrap("first_function")
def first_function():
    second_function()

@tracer.wrap("second_function")
def second_function():
    try:
        third_function()
    except ValueError as e:
        raise ValueError('error when calling third_function in second_function') from e

@tracer.wrap("third_function")
def third_function():
    try:
        fourth_function()
    except ValueError as e:
        raise ValueError('error when calling fourth_function in third_function')

@tracer.wrap("fourth_function")
def fourth_function():
    try:
        s = "baba"
        a = s.split("a")
        print(a[5])
    except IndexError as e:
        raise ValueError('failed in fourth_function')
