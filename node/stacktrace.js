function trigger_function() {
  try {
    throw new Error("manually generated exception");
  } catch (e) {
    throw new Error("failed in trigger_function");
  }
}

function second_function() {
  try {
    trigger_function();
  } catch (e) {
    throw new Error("failed in second_function");
  }
}

function first_function() {
  second_function();
}

first_function();
