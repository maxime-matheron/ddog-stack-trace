<?php
function trigger_exception() {
    try {
        throw new InvalidArgumentException('manually generated exception');
    } catch (Exception $e) {
        throw new LogicException('failed in trigger_exception', 0, $e);
    }
}

function second_function() {
    try {
        trigger_exception();
    } catch (Exception $e) {
        throw new LogicException('failed in second_function', 0, $e);
    }
}

function first_function() {
    second_function();
}

first_function()
?>


