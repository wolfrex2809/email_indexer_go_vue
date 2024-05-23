<script setup lang="ts">
import instance from './utils/axios';
import { onMounted, ref } from 'vue';
import { Email } from './models/email.model';

const emails = ref<Email[]>([]);
const getEmails = async () => {
  try{
    const { data } = await instance.get("/emails/search", {
      params: {
        text: "a"
      }
    })
    emails.value = data.data;
  } catch(error){
    console.log(error)
  }

}

onMounted(() => {
  getEmails()
})
</script>

<template>
  <div class="text-center mt-8">
    <img src="/logo.png" class="mx-auto" alt="logo" width="100" />
    <h1 class="mt-2 text-2xl text-blue-400 font-bold font-mono">Email Searcher</h1>
  </div>

  <div class="max-w-2xl mx-auto mt-6 px-3">
    <form>   
        <label for="default-search" class="mb-2 text-sm font-medium text-gray-900 sr-only dark:text-gray-300">Search</label>
        <div class="relative">
            <div class="flex absolute inset-y-0 left-0 items-center pl-3 pointer-events-none">
                <svg class="w-5 h-5 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
            </div>
            <input type="search" id="default-search" class="block p-4 pl-10 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-800 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Search Emails" required>
            <button type="submit" class="absolute outline right-2.5 bottom-2.5 focus:ring-4 hover:outline-blue-400  font-medium rounded-lg text-sm px-4 py-2">Search</button>
        </div>
    </form>
  </div>
  <div class="mt-6 grid grid-cols-12 outline outline-blue-400 min-h-96 max-h-96 mx-2 rounded-md overflow-y-hidden">
    <div class="col-span-3 overflow-y-auto bg-gray-800 max-h-96 divide-y">
      <div v-if="emails == null">
        <h1 class="text-gray-400 text-center mt-5">No results</h1>
      </div>
      <div v-for="email in emails" v-if="emails != null" class="p-3 hover:bg-gray-900">
        <a href="#">
          <div>
            <h5 class="flex">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-5 text-blue-400">
                <path stroke-linecap="round" stroke-linejoin="round" d="M21.75 6.75v10.5a2.25 2.25 0 0 1-2.25 2.25h-15a2.25 2.25 0 0 1-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0 0 19.5 4.5h-15a2.25 2.25 0 0 0-2.25 2.25m19.5 0v.243a2.25 2.25 0 0 1-1.07 1.916l-7.5 4.615a2.25 2.25 0 0 1-2.36 0L3.32 8.91a2.25 2.25 0 0 1-1.07-1.916V6.75" />
              </svg>
              <p class="truncate ml-2">{{ email.subject }}</p></h5>
            <span class="flex">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-5 text-blue-400">
                <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 6a3.75 3.75 0 1 1-7.5 0 3.75 3.75 0 0 1 7.5 0ZM4.501 20.118a7.5 7.5 0 0 1 14.998 0A17.933 17.933 0 0 1 12 21.75c-2.676 0-5.216-.584-7.499-1.632Z" />
              </svg>
              <span class="text-xs ml-2">{{ email.from }}</span></span>
            <p class="truncate text-gray-500 text-sm">{{ email.content }}</p>
          </div>
        </a>
      </div>
    </div>
    <div class="col-span-9 overflow-y-auto bg-gray-800 border-l-4 border-blue-400 p-3">
      <!-- <h1 class="text-center text-4xl text-gray-500">Nothing to show</h1> Default view-->
      <div>
        <h1 class="text-2xl">Hi this is a long subject for the email searcher</h1 class="text-2xl">
        <span class="text-gray-400">From: <span class="text-sm text-white">xxxxxxxxxx@xxx.com</span></span>
        <br>
        <span class="text-gray-400">To: <span class="text-sm text-white">xxxxxxxxxx@xxx.com, xxxxxxxxxx2@xxx.com, xxxxxxxxxx3@xxx.com</span></span>
        <br>
        <span class="text-gray-400">Date: <span class="text-sm text-white">23-05-2024 00:00:00 (GMT -4:00)</span></span>
      </div>
      <div class="mt-6">
        <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Dolorum saepe repellendus aperiam facere illum quas ut! Earum fuga vel dignissimos voluptatem ad rerum consequatur quas? Praesentium ipsa iusto aspernatur vel.</p>
      </div>
    </div>
  </div>


</template>
