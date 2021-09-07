from ddtrace import tracer

@tracer.wrap("call.first_function")
def first_function():
    second_function()

@tracer.wrap("call.second_function")
def second_function():
    try:
        third_function()
    except ValueError as e:
        raise ValueError('failed to call third_function in second_function') from e

@tracer.wrap("call.third_function")
def third_function():
    try:
        fourth_function()
    except ValueError as e:
        raise ValueError('failed to call fourth_function in third_function')

@tracer.wrap("call.fourth_function")
def fourth_function():
    try:
        s = "baba"
        a = s.split("a")
        print(a[5])
    except IndexError as e:
        raise ValueError('failed in fourth_function')
