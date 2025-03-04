import { writable } from 'svelte/store'

export const uploadHistoryState = writable(false)
export const deleteAccountState = writable(false)
export const errorState = writable({
  title: '',
  message: ''
})
export const loadingState = writable(false)