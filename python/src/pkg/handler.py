from ddtrace import tracer

def first_function():
    with tracer.trace('call.second_function'):
        second_function()

def second_function():
    try:
        with tracer.trace('call.third_function'):
            third_function()
    except ValueError as e:
        raise ValueError('failed to call third_function in second_function') from e

def third_function():
    try:
        with tracer.trace('call.fourth_function'):
            fourth_function()
    except ValueError as e:
        raise ValueError('failed to call fourth_function in third_function')

def fourth_function():
    try:
        with tracer.trace('run.operation'):
            v = {}['a']
    except KeyError as e:
        raise ValueError('failed in fourth_function')
