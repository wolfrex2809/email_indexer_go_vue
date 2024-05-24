<script setup lang="ts">
import instance from './utils/axios';
import { onMounted, ref } from 'vue';
import { Email } from './models/email.model';
import CurrentEmail from "./components/CurrentEmail.vue"
import EmailList from "./components/EmailList.vue"

const emails = ref<Email[]>();
const text = ref<string>("");
var currentEmail = ref<Email>();
const CurrentEmailContainer = ref();
const EmailListContainer = ref();

const getEmails = async (text: string = "a") => {
  try{
    const { data } = await instance.get("/emails/search", {
      params: {
        text: text
      }
    })
    emails.value = data.data;
    EmailListContainer.value.scrollTop = 0;
    currentEmail = ref<Email>()
  } catch(error){
    console.log(error)
  }

}

const submit = async () => {
  try{
    getEmails(text.value)
  } catch(error){
    console.log(error)
  }
}

const viewEmail = (email: Email) => {
  try{
    currentEmail.value = email; 
    CurrentEmailContainer.value.scrollTop = 0;
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
    <form @submit.prevent="submit">   
        <label for="default-search" class="mb-2 text-sm font-medium text-gray-900 sr-only dark:text-gray-300">Search</label>
        <div class="relative">
            <div class="flex absolute inset-y-0 left-0 items-center pl-3 pointer-events-none">
                <svg class="w-5 h-5 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
            </div>
            <input type="search" id="default-search" v-model="text" class="block p-4 pl-10 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-800 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Search Emails" required>
        </div>
    </form>
  </div>
  <div class="mt-6 grid grid-cols-12 outline outline-blue-400 min-h-96 max-h-96 mx-2 rounded-md overflow-y-hidden">
    <div class="col-span-3 overflow-y-auto bg-gray-800 max-h-96 divide-y" ref="EmailListContainer">
      <EmailList :emails="emails" :currentEmail="currentEmail" @select-email="viewEmail"/>
    </div>
    <div class="col-span-9 overflow-y-auto bg-gray-800 border-l-4 border-blue-400 p-3 max-h-96" ref="CurrentEmailContainer">
      <CurrentEmail :currentEmail="currentEmail"/>

    </div>
  </div>


</template>
