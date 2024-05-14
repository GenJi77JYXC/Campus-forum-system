import {defineStore} from 'pinia'
import type { UserState } from './modules/type'
import { GET_TOKEN } from '@/utils/token'


let useUserStore = defineStore('User', {
    state: ():UserState => {
        return {
            token: GET_TOKEN(),
            username: ''
        }
    }
})