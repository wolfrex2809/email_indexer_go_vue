<script setup lang="ts">
import { Email } from '../models/email.model.ts';

interface Props {
  currentEmail: Email | undefined
}

defineProps<Props>();

const parseDate = (date:string) => {
    return new Date(date).toLocaleString()
}

</script>
<template>
    <h1 class="text-center text-4xl text-gray-500" v-if="currentEmail == null">Nothing to show</h1>
    <div v-if="currentEmail != null">
        <div>
          <h1 class="text-2xl">{{ currentEmail.subject }}</h1 class="text-2xl">
          <span class="text-gray-400">From: <span class="text-sm text-white">{{ currentEmail.from != "" ? currentEmail.from : "N/A" }}</span></span>
          <br>
          <span class="text-gray-400">To: <span class="text-sm text-white">{{ currentEmail.to != "" ? currentEmail.to : "N/A" }}</span></span>
          <br>
          <span class="text-gray-400">Date: <span class="text-sm text-white">{{ parseDate(currentEmail.date) }}</span></span>
        </div>
        <div class="mt-6">
          <p class="whitespace-pre">{{ currentEmail.content }}</p>
        </div>
    </div>
</template>        