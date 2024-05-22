<script setup lang="ts">
import HelloWorld from './components/HelloWorld.vue'
</script>

<template>
  <div class="text-center mt-8">
    <img src="/logo.png" class="mx-auto" alt="logo" width="100" />
    <h1 class="mt-2 text-2xl">Email Searcher</h1>
  </div>

  <div class="max-w-2xl mx-auto mt-6 px-3">
    <form>   
        <label for="default-search" class="mb-2 text-sm font-medium text-gray-900 sr-only dark:text-gray-300">Search</label>
        <div class="relative">
            <div class="flex absolute inset-y-0 left-0 items-center pl-3 pointer-events-none">
                <svg class="w-5 h-5 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
            </div>
            <input type="search" id="default-search" class="block p-4 pl-10 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Search Emails" required>
            <button type="submit" class="absolute outline right-2.5 bottom-2.5 focus:ring-4 hover:outline-blue-400  font-medium rounded-lg text-sm px-4 py-2">Search</button>
        </div>
    </form>
  </div>
  <div class="mt-6 grid grid-cols-12 outline min-h-96 max-h-96 mx-2 rounded-md overflow-y-hidden">
    <div class="col-span-3 overflow-y-auto bg-gray-800 max-h-96 divide-y">
      <div class="p-3 hover:bg-gray-900">
        <a href="#">
          <div>
            <h5 class="truncate">Hi this is a long subject for the email searcher</h5>
            <span class="text-xs">From: <span class="xxxxxxxxxx@xxx.com">xxxxxxxxxx@xxx.com</span></span>
            <p class="truncate text-gray-500 text-sm">Lorem ipsum dolor sit amet consectetur adipisicing elit. Corporis totam tempora atque deleniti, iste saepe! Obcaecati repellendus dolores doloribus fuga, accusamus perspiciatis laborum quia cumque nemo. Nesciunt quia delectus rem!</p>
          </div>
        </a>
      </div>
      <div class="p-3 hover:bg-gray-900">
        <a href="#">
          <div>
            <h5 class="truncate">Hi this is a long subject for the email searcher</h5>
            <span class="text-xs">From: <span class="xxxxxxxxxx@xxx.com">xxxxxxxxxx@xxx.com</span></span>
            <p class="truncate text-gray-500 text-sm">Lorem ipsum dolor sit amet consectetur adipisicing elit. Corporis totam tempora atque deleniti, iste saepe! Obcaecati repellendus dolores doloribus fuga, accusamus perspiciatis laborum quia cumque nemo. Nesciunt quia delectus rem!</p>
          </div>
        </a>
      </div>
      <div class="p-3 hover:bg-gray-900">
        <a href="#">
          <div>
            <h5 class="truncate">Hi this is a long subject for the email searcher</h5>
            <span class="text-xs">From: <span class="xxxxxxxxxx@xxx.com">xxxxxxxxxx@xxx.com</span></span>
            <p class="truncate text-gray-500 text-sm">Lorem ipsum dolor sit amet consectetur adipisicing elit. Corporis totam tempora atque deleniti, iste saepe! Obcaecati repellendus dolores doloribus fuga, accusamus perspiciatis laborum quia cumque nemo. Nesciunt quia delectus rem!</p>
          </div>
        </a>
      </div>
      <div class="p-3 hover:bg-gray-900">
        <a href="#">
          <div>
            <h5 class="truncate">Hi this is a long subject for the email searcher</h5>
            <span class="text-xs">From: <span class="xxxxxxxxxx@xxx.com">xxxxxxxxxx@xxx.com</span></span>
            <p class="truncate text-gray-500 text-sm">Lorem ipsum dolor sit amet consectetur adipisicing elit. Corporis totam tempora atque deleniti, iste saepe! Obcaecati repellendus dolores doloribus fuga, accusamus perspiciatis laborum quia cumque nemo. Nesciunt quia delectus rem!</p>
          </div>
        </a>
      </div>
      <div class="p-3 hover:bg-gray-900">
        <a href="#">
          <div>
            <h5 class="truncate">Hi this is a long subject for the email searcher</h5>
            <span class="text-xs">From: <span class="xxxxxxxxxx@xxx.com">xxxxxxxxxx@xxx.com</span></span>
            <p class="truncate text-gray-500 text-sm">Lorem ipsum dolor sit amet consectetur adipisicing elit. Corporis totam tempora atque deleniti, iste saepe! Obcaecati repellendus dolores doloribus fuga, accusamus perspiciatis laborum quia cumque nemo. Nesciunt quia delectus rem!</p>
          </div>
        </a>
      </div>
      <div class="p-3 hover:bg-gray-900">
        <a href="#">
          <div>
            <h5 class="truncate">Hi this is a long subject for the email searcher</h5>
            <span class="text-xs">From: <span class="xxxxxxxxxx@xxx.com">xxxxxxxxxx@xxx.com</span></span>
            <p class="truncate text-gray-500 text-sm">Lorem ipsum dolor sit amet consectetur adipisicing elit. Corporis totam tempora atque deleniti, iste saepe! Obcaecati repellendus dolores doloribus fuga, accusamus perspiciatis laborum quia cumque nemo. Nesciunt quia delectus rem!</p>
          </div>
        </a>
      </div>
      <div class="p-3 hover:bg-gray-900">
        <a href="#">
          <div>
            <h5 class="truncate">Hi this is a long subject for the email searcher</h5>
            <span class="text-xs">From: <span class="xxxxxxxxxx@xxx.com">xxxxxxxxxx@xxx.com</span></span>
            <p class="truncate text-gray-500 text-sm">Lorem ipsum dolor sit amet consectetur adipisicing elit. Corporis totam tempora atque deleniti, iste saepe! Obcaecati repellendus dolores doloribus fuga, accusamus perspiciatis laborum quia cumque nemo. Nesciunt quia delectus rem!</p>
          </div>
        </a>
      </div>
      <div class="p-3 hover:bg-gray-900">
        <a href="#">
          <div>
            <h5 class="truncate">Hi this is a long subject for the email searcher</h5>
            <span class="text-xs">From: <span class="xxxxxxxxxx@xxx.com">xxxxxxxxxx@xxx.com</span></span>
            <p class="truncate text-gray-500 text-sm">Lorem ipsum dolor sit amet consectetur adipisicing elit. Corporis totam tempora atque deleniti, iste saepe! Obcaecati repellendus dolores doloribus fuga, accusamus perspiciatis laborum quia cumque nemo. Nesciunt quia delectus rem!</p>
          </div>
        </a>
      </div>
    </div>
    <div class="col-span-9 overflow-y-auto bg-gray-800 border-l-4 text-center">
      <h1 class="text-4xl">---Here is the email content---</h1>
    </div>
  </div>


</template>
