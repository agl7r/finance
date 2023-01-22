<script setup>
import { ref } from 'vue'
import menu from './menu.js'
import { useRoute } from 'vue-router'

let leftDrawerOpen = ref(false)

const toggleLeftDrawer = function() {
  leftDrawerOpen.value = !leftDrawerOpen.value
}

const route = useRoute()

const isCurrentPath = (menuItem) => {
  return menuItem.name === route.name
}

</script>

<template>
  <q-layout view="hHh lpR fFf">

    <q-header elevated class="bg-primary text-white">
      <q-toolbar>
        <q-btn dense flat round icon="menu" @click="toggleLeftDrawer" />

        <q-toolbar-title>
          Финансы
        </q-toolbar-title>
      </q-toolbar>
    </q-header>

    <q-drawer show-if-above v-model="leftDrawerOpen" side="left" bordered>
      <q-list>
        <template v-for="(menuItem, index) in menu" :key="index">
          <q-item clickable :to="{ name: menuItem.name }" :active="isCurrentPath(menuItem)" v-ripple>
            <q-item-section>
              {{ menuItem.label }}
            </q-item-section>
          </q-item>
        </template>
      </q-list>
    </q-drawer>

    <q-page-container>
      <q-page padding>
        <router-view></router-view>
      </q-page>
    </q-page-container>

  </q-layout>
</template>

<style>
h1 {
  font-size: 2rem;
  line-height: 2rem;
}
</style>
