<script setup lang="ts">
import { Email } from '../models/email.model.ts';

interface Props {
    emails: Email[] | undefined
    currentEmail: Email | undefined
}

defineProps<Props>();
const emit = defineEmits(['selectEmail'])

const parseDate = (date:string) => {
    var formated = new Date(date).toLocaleString()
    return formated.split(",")[0]
}
</script>
<template>
    <div v-if="emails == null">
        <h1 class="text-gray-400 text-center mt-5">No results</h1>
    </div>
    <div v-for="email in emails" v-if="emails != null" class="p-3 hover:bg-gray-900" :class="{'bg-gray-900': email == currentEmail}">
        <div @click="$emit('selectEmail', email)" class="cursor-pointer">
            <div class="grid grid-cols-12">
                <h5 class="flex col-span-9">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-5 text-blue-400">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M21.75 6.75v10.5a2.25 2.25 0 0 1-2.25 2.25h-15a2.25 2.25 0 0 1-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0 0 19.5 4.5h-15a2.25 2.25 0 0 0-2.25 2.25m19.5 0v.243a2.25 2.25 0 0 1-1.07 1.916l-7.5 4.615a2.25 2.25 0 0 1-2.36 0L3.32 8.91a2.25 2.25 0 0 1-1.07-1.916V6.75" />
                    </svg>
                    <p class="truncate ml-2">{{ email.subject }}</p>
                </h5>
                <span class="text-xs col-span-3 text-end">
                    {{ parseDate(email.date) }}
                </span>
            </div>
            <span class="flex">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="size-5 text-blue-400">
                    <path fill-rule="evenodd" d="M12.97 3.97a.75.75 0 0 1 1.06 0l7.5 7.5a.75.75 0 0 1 0 1.06l-7.5 7.5a.75.75 0 1 1-1.06-1.06l6.22-6.22H3a.75.75 0 0 1 0-1.5h16.19l-6.22-6.22a.75.75 0 0 1 0-1.06Z" clip-rule="evenodd" />
                </svg>
                <span class="text-xs ml-2">{{ email.from != "" ? email.from : "N/A" }}</span>
            </span>
            <span class="flex">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="size-5 text-blue-400">
                    <path fill-rule="evenodd" d="M11.03 3.97a.75.75 0 0 1 0 1.06l-6.22 6.22H21a.75.75 0 0 1 0 1.5H4.81l6.22 6.22a.75.75 0 1 1-1.06 1.06l-7.5-7.5a.75.75 0 0 1 0-1.06l7.5-7.5a.75.75 0 0 1 1.06 0Z" clip-rule="evenodd" />
                </svg>
                <span class="text-xs ml-2">{{ email.to != "" ? email.to : "N/A" }}</span>
            </span>
            <p class="truncate text-gray-500 text-sm">{{ email.content }}</p>
        </div>
    </div>
</template>     