<script setup lang="ts">
import { Email } from '../models/email.model.ts';

interface Props {
    emails: Email[] | undefined
    currentEmail: Email | undefined
}

defineProps<Props>();
const emit = defineEmits(['selectEmail'])

</script>
<template>
    <div v-if="emails == null">
        <h1 class="text-gray-400 text-center mt-5">No results</h1>
    </div>
    <div v-for="email in emails" v-if="emails != null" class="p-3 hover:bg-gray-900" :class="{'bg-gray-900': email == currentEmail}">
        <div @click="$emit('selectEmail', email)" class="cursor-pointer">
            <h5 class="flex">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-5 text-blue-400">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M21.75 6.75v10.5a2.25 2.25 0 0 1-2.25 2.25h-15a2.25 2.25 0 0 1-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0 0 19.5 4.5h-15a2.25 2.25 0 0 0-2.25 2.25m19.5 0v.243a2.25 2.25 0 0 1-1.07 1.916l-7.5 4.615a2.25 2.25 0 0 1-2.36 0L3.32 8.91a2.25 2.25 0 0 1-1.07-1.916V6.75" />
                </svg>
                <p class="truncate ml-2">{{ email.subject }}</p>
            </h5>
            <span class="flex">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-5 text-blue-400">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 6a3.75 3.75 0 1 1-7.5 0 3.75 3.75 0 0 1 7.5 0ZM4.501 20.118a7.5 7.5 0 0 1 14.998 0A17.933 17.933 0 0 1 12 21.75c-2.676 0-5.216-.584-7.499-1.632Z" />
                </svg>
                <span class="text-xs ml-2">{{ email.from }}</span>
            </span>
            <p class="truncate text-gray-500 text-sm">{{ email.content }}</p>
        </div>
    </div>
</template>     