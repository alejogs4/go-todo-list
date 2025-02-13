function wrappedFetch(url, options = {}) {
  return fetch(`http://localhost:8080${url}`, {
    ...options,
    headers: {
      "Content-Type": "application/json",
      ...options.headers || {},
    },
    body: options.body ? JSON.stringify(options.body) : undefined,
  }).then((response) => {
    if (!response.ok) {
      throw new Error(response.statusText);
    }

    if (response.status >= 400) {
      throw new Error(response.status);
    }

    if (response.status === 204) {
      return;
    }

    return response.json();
  });
}

const todoService = {
  getAllTodos: () => wrappedFetch("/api/v1/todo"),
  createTodo: (todo) => wrappedFetch("/api/v1/todo", { method: "POST", body: todo }),
  updateTodo: (todoID, todo) => wrappedFetch(`/api/v1/todo/${todoID}`, { method: "PUT", body: todo }),
  deleteTodo: (todoID) => wrappedFetch(`/api/v1/todo/${todoID}`, { method: "DELETE" }),
};


async function displayTodos() {
  const todosContainer = document.getElementById("todo_list");
  if (!todosContainer) {
    return;
  }

  const todos = await todoService.getAllTodos();

  todosContainer.innerHTML = "";
  todos.content.forEach((todo) => {
    const todoElement = document.createElement("li");
    todoElement.className = "todo-item";
    const contentContainer = document.createElement("div");

    const todoRemoveButton = document.createElement("button");
    todoRemoveButton.innerText = "X";
    todoRemoveButton.addEventListener("click", removeTodo(todo));

    const todoCompleteCheckbox = document.createElement("input");
    todoCompleteCheckbox.type = "checkbox";
    todoCompleteCheckbox.checked = todo.completed;
    todoCompleteCheckbox.addEventListener("change", completeTodo(todo));

    contentContainer.innerText = todo.title;
    contentContainer.contentEditable = true;
    contentContainer.className = "todo-content";

    todoElement.appendChild(contentContainer);
    todoElement.appendChild(todoCompleteCheckbox);
    todoElement.appendChild(todoRemoveButton);

    contentContainer.addEventListener("blur", editTodo(todo));

    todosContainer.appendChild(todoElement);
  });
}

const todoAuditService = {
  getAllAudits: () => wrappedFetch("/api/v1/report"),
}

async function displayTodosReport() {
  const todosReportContainer = document.getElementById("report_list");
  if (!todosReportContainer) {
    return;
  }

  const todoReport = await todoAuditService.getAllAudits()
  
  todosReportContainer.innerHTML = "";
  todoReport.content.forEach((todoReport) => {
    const todoReportElement = document.createElement("li");
    todoReportElement.innerText = todoReport.description;
    todosReportContainer.appendChild(todoReportElement);
  });
}

function editTodo(todo) {
  return async e => {
    todo.title = e.target.innerText;
    await todoService.updateTodo(todo.id, { title: todo.title.replaceAll("\n", ''), completed: todo.completed });
    await displayTodosReport();
  }
}

function removeTodo(todo) {
  return async e => {
    await todoService.deleteTodo(todo.id)
    e.target.parentElement.remove();
    await Promise.all([displayTodosReport(), displayTodos()]);
  }
}

function completeTodo(todo) {
  return async e => {
    await todoService.updateTodo(todo.id, { title: todo.title, completed: e.target.checked });
    await displayTodosReport();
  }
}

const todoForm = document.getElementById("todo_form");
async function createTodo(e) {
  e.preventDefault();
  const todo = { title: e.target.todo_input.value,  };
  
  await todoService.createTodo(todo);
  e.target.todo_input.value = "";
  await Promise.all([displayTodosReport(), displayTodos()]);
}
todoForm.addEventListener("submit", createTodo);

displayTodos();
displayTodosReport();
