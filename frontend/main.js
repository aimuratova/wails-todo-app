// frontend/main.js


// `backend` functions are available because мы привязали TodoService в Go

import { AddTask, GetTasks, ToggleDone, UpdateTask, DeleteTask } 
  from "./wailsjs/go/backend/TodoService";

async function el(id){ return document.getElementById(id); }

async function renderTasks() {
    const tasks = await GetTasks();
    const list = document.getElementById('tasks');
    list.innerHTML = '';
    tasks.forEach(task => {
    const li = document.createElement('li');
    li.className = task.done ? 'done' : '';


    const checkbox = document.createElement('input');
    checkbox.type = 'checkbox';
    checkbox.checked = task.done;
    checkbox.addEventListener('change', async () => {
        await ToggleDone(task.id);
        await renderTasks();
    });


    const span = document.createElement('div');
    span.className = 'task-text';
    span.textContent = task.text;


    const actions = document.createElement('div');
    actions.className = 'task-actions';


    const editBtn = document.createElement('button');
    editBtn.textContent = 'Изменить';
    editBtn.addEventListener('click', async () => {
        const newText = prompt('Изменить задачу', task.text);
        if (newText !== null) {
            await UpdateTask(task.id, newText);
            await renderTasks();
    }
    });


    const delBtn = document.createElement('button');
    delBtn.textContent = 'Удалить';
    delBtn.addEventListener('click', async () => {
        if (confirm('Удалить задачу?')){
            await DeleteTask(task.id);
            await renderTasks();
        }
    });


    actions.appendChild(editBtn);
    actions.appendChild(delBtn);


    li.appendChild(checkbox);
    li.appendChild(span);
    li.appendChild(actions);


    list.appendChild(li);
    });
}

async function init() {
    document.getElementById('addBtn').addEventListener('click', async () => {
        const input = document.getElementById('newTaskInput');
        const text = input.value.trim();
        if (!text) return;
        await AddTask(text);
        input.value = '';
        await renderTasks();
    });


    // Enter key
    document.getElementById('newTaskInput').addEventListener('keydown', async (e)=>{
        if (e.key === 'Enter') document.getElementById('addBtn').click();
    });


    await renderTasks();
}


window.addEventListener('DOMContentLoaded', init);