<script setup lang="ts">
import instance from './utils/axios';
import { onMounted, ref } from 'vue';
import { Email } from './models/email.model';
import CurrentEmail from "./components/CurrentEmail.vue"
import EmailList from "./components/EmailList.vue"

const emails = ref<Email[]>();
const text = ref<string>("");
const currentEmail = ref<Email>();
const CurrentEmailContainer = ref();
const EmailListContainer = ref();
const isLoading = ref<boolean>(false);
var total = 0;
var page = 0;


const getEmails = async (text: string = "a", page_num: number = 0) => {
  try{
    isLoading.value = true;
    console.log(isLoading)
    currentEmail.value = undefined
    const { data } = await instance.get("/emails/search", {
      params: {
        text: text,
        page: page_num
      }
    })
    isLoading.value = false;
    console.log(isLoading)
    emails.value = data.data;
    total = data.total;
    EmailListContainer.value.scrollTop = 0;
  } catch(error){
    isLoading.value = false;
    console.log(error)
  }

}

const submit = async () => {
  try{
    total = 0;
    page = 0;
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

const clearView = () => {
  try{
    currentEmail.value = undefined
  } catch(error){
    console.log(error)
  }
}

const isFirstPage = () => {
  return page == 0
}

const isLastPage = () => {
  var result = (total / 10) - (page + 1);
  return result < 0;
}

const searchNext = () => {
  page++;
  getEmails(text.value ? text.value : "a", page)
}
const searchPrev = () => {
  page--;
  getEmails(text.value ? text.value : "a", page)
}

onMounted(() => {
  getEmails()
})

</script>

<template>
  <div v-if="currentEmail" @click="clearView" class="w-8 z-10 max-sm:block md:hidden top-2 start-2 absolute">
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="size-6">
      <path fillRule="evenodd" d="M11.03 3.97a.75.75 0 0 1 0 1.06l-6.22 6.22H21a.75.75 0 0 1 0 1.5H4.81l6.22 6.22a.75.75 0 1 1-1.06 1.06l-7.5-7.5a.75.75 0 0 1 0-1.06l7.5-7.5a.75.75 0 0 1 1.06 0Z" clipRule="evenodd" />
    </svg>
  </div>
  <div class="text-center mt-7">
    <img src="/logo.png" class="mx-auto" alt="logo" width="100" />
    <h1 class="mt-2 text-2xl text-blue-400 font-bold font-mono">Email Searcher</h1>
  </div>

  <div class="max-w-2xl mx-auto md:mt-5 px-3">
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
  <div class="max-sm:hidden md:grid mt-6 grid grid-cols-12 outline outline-blue-400 min-h-96 max-h-96 mx-3 rounded-md overflow-y-hidden bg-gray-800">
    <div class="col-span-3 bg-gray-800">
      <div class="h-80 flex items-center justify-center" v-if="isLoading">
        <div role="status" class="text-center">
            <svg aria-hidden="true" class="w-8 h-8 text-gray-200 animate-spin dark:text-gray-600 fill-blue-600" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="currentColor"/>
                <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentFill"/>
            </svg>
            <span class="sr-only">Loading...</span>
        </div>
      </div>
      <div class="overflow-y-auto max-h-96 divide-y" v-if="!isLoading" ref="EmailListContainer">
        <EmailList :emails="emails" :currentEmail="currentEmail" @select-email="viewEmail"/>
      </div>
    </div>
    <div class="col-span-9 overflow-y-auto bg-gray-800 border-l-4 border-blue-400 p-3 max-h-96" ref="CurrentEmailContainer">
      <CurrentEmail :currentEmail="currentEmail"/>
    </div>
  </div>
  <div class="max-sm:block md:hidden mt-6 grid grid-cols-12 outline outline-blue-400 min-h-96 max-h-96 mx-3 rounded-md overflow-y-hidden bg-gray-800">
    <div v-if="!currentEmail" class="bg-gray-800">
      <div class="h-80 flex items-center justify-center" v-if="isLoading">
        <div role="status" class="text-center">
            <svg aria-hidden="true" class="w-8 h-8 text-gray-200 animate-spin dark:text-gray-600 fill-blue-600" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="currentColor"/>
                <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentFill"/>
            </svg>
            <span class="sr-only">Loading...</span>
        </div>
      </div>
      <div class="overflow-y-auto max-h-96 divide-y" v-if="!isLoading" ref="EmailListContainer">
        <EmailList :emails="emails" :currentEmail="currentEmail" @select-email="viewEmail"/>
      </div>
    </div>
    <div v-if="currentEmail" class="overflow-y-auto bg-gray-800 p-3 max-h-96" ref="CurrentEmailContainer">
      <CurrentEmail :currentEmail="currentEmail"/>
    </div>
  </div>

  <div class="m-3 grid-cols-12 md:grid">
    <div class="col-span-3 grid grid-cols-12">
      <div class="col-span-8">
        <h5>{{ total != 0 ? 10 * page + 1 : 0 }}-{{ 10 * (page + 1) > total ? total : 10 * (page + 1) }} of {{ total }} results</h5>
      </div>
      <div class="col-span-4 text-right">
        <div class="flex flex-row-reverse h-8 text-base">
          <button :disabled="isLastPage() || isLoading" v-on:click="searchNext" class="flex items-center justify-center px-4 h-8 leading-tight disabled:opacity-75 text-gray-500 bg-white border border-gray-300 rounded-e-lg enabled:hover:bg-gray-100 enabled:hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 enabled:dark:hover:bg-gray-700 enabled:dark:hover:text-white">
            <span class="sr-only">Next</span>
            <svg class="w-3 h-3 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 6 10">
              <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 9 4-4-4-4"/>
            </svg>
          </button>
          <button :disabled="isFirstPage() || isLoading" v-on:click="searchPrev" class="flex items-center justify-center px-4 h-8 ms-0 leading-tight disabled:opacity-75 text-gray-500 bg-white border border-e-0 border-gray-300 rounded-s-lg enabled:hover:bg-gray-100 enabled:hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 enabled:dark:hover:bg-gray-700 enabled:dark:hover:text-white">
            <span class="sr-only">Previous</span>
            <svg class="w-3 h-3 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 6 10">
              <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 1 1 5l4 4"/>
            </svg>
          </button>
        </div>
      </div>
    </div>
  </div>


</template>
