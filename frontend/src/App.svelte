<script lang="ts">
  import Todo from "./lib/Todo.svelte";
  import type { TodoItem } from "./lib/types";

  let todos: TodoItem[] = $state([]);

  async function fetchTodos() {
    try {

      // sends a GET request to the backend server
      const response = await fetch("http://localhost:8080/");
      if (response.status !== 200) {
        console.error("Error fetching data. Response status not 200");
        return;
      }

      todos = await response.json();
    } catch (e) {
      console.error("Could not connect to server. Ensure it is running.", e);
    }
  }

  //when add todo button is pressed, this function should be called. It seems to innately contain the event parameter despite being called without arguments.
  async function handleAddTodo(event: Event) {

    event.preventDefault(); // I had to add this to prevent the page from reloading on form submission which was breaking functionality

    // get the form that triggered the event
    const form = event.target as HTMLFormElement; // I needed to explicitly define the type here to avoid an error just below this

    // extract all the form data
    const data = new FormData(form); // form was throwing an error here, so I explicitly defined its type above

    // get each value by its field name
    const title = data.get("title");
    const description = data.get("description");

    // after parsing data, send the request to the backend. Using a try to stop errors from causing crashes
    try {

      //send the POST request to the server at the correct port 8080
      const response = await fetch("http://localhost:8080/", {

        method: "POST",

        //specifies JSON content using a standard header
        headers: { "Content-Type": "application/json" },

        //attempt to build the JSON and send it, reading in the following format:
        // send this data:  JSON object
        body: JSON.stringify({ title, description })
      });

      // log an error in the console if the response is anything other than 200 (200 means success)
      if (response.status !== 200) {
        console.error("Failed to add todo.");
        return;
      }

      // the backend should return the newly created todo item as confirmation, if successful, as the spec outlines
      const newTodo = await response.json();

      //  add the newly created item to the todos array to update the UI
      //"..." operator means a new array is created instead of updating the existing one
      todos = [...todos, newTodo];

    
    //if the try fails, then it is most likely no server was found, so throw error that no server could be connected to.
    } catch (e) {
      console.error("Could not connect to server to add todo.", e);
    }

  }

  // Initially fetch todos on page load
  $effect(() => {
    fetchTodos();
  });
</script>

<main class="app">
  <header class="app-header">
    <h1>TODO</h1>
  </header>

  <div class="todo-list">
    {#each todos as todo}
      <Todo title={todo.title} description={todo.description} />
    {/each}
  </div>

  <h2 class="todo-list-form-header">Add a Todo</h2>

  <!-- added a call to handleAddTodo when the form is submitted -->
  <form class="todo-list-form">
    <input placeholder="Title" name="title" />
    <input placeholder="Description" name="description" />
    <button>Add Todo</button>
  </form>
</main>

<style>
  .app {
    color: white;
    background-color: #282c34;

    text-align: center;
    font-size: 24px;

    min-height: 100vh;
    padding: 20px;
  }

  .app-header {
    font-size: calc(10px + 4vmin);
    margin-top: 50px;
  }

  .todo-list {
    margin: 50px 100px 0px 100px;
  }

  .todo-list-form-header {
    margin-top: 100px;
  }

  .todo-list-form {
    margin-top: 10px;
  }
</style>
