<script setup>
import { computed, ref, watch } from 'vue'
import { fetchUtilityBillsByYears } from '../../api/utilityBills.js'

const getTypes = function (bills) {
  const types = []
  bills.forEach(bill => {
    const foundType = types.find(m => m.id === bill.type.id)
    if (!foundType) {
      types.push(bill.type)
    }
  })
  types.sort((a, b) => a.id > b.id ? 1 : -1)
  return types
}

const buildColumns = function (bills) {
  const types = getTypes(bills)

  return [
    { label: 'Месяц', field: 'month', align: 'left' },
    ...types.map(type => {
      return { label: type.title, field: type.id, align: 'left' }
    })
  ]
}

const getMonths = function (bills) {
  const months = []
  bills.forEach(bill => {
    const foundMonth = months.find(month => month.id === bill.month.id)
    if (!foundMonth) {
      months.push(bill.month)
    }
  })
  months.sort((a, b) => a.id > b.id ? 1 : -1)
  return months
}

const buildRows = function (bills) {
  const rows = []
  const months = getMonths(bills)
  months.forEach(month => {
    const monthBills = bills.filter(payment => payment.month.id === month.id)
    if (monthBills.length > 0) {
      const row = { month: month.id }
      monthBills.forEach(bill => {
        row[bill.type.id] = bill.amount.number
      })
      rows.push(row)
    }
  })
  return rows
}

let year = ref((new Date()).getFullYear())

let pageTitle = computed(() => `Коммунальные платежи за ${year.value} год`)
let loading = ref(true)
let columns = ref([])
let rows = ref([])

const fetchData = (year) => {
  fetchUtilityBillsByYears(year).then(bills => {
    columns.value = buildColumns(bills)
    rows.value = buildRows(bills)
    loading.value = false
  })
}

fetchData(year.value)
watch(year, () => { fetchData(year.value) })

</script>

<template>
  <h1>{{ pageTitle }}</h1>

  <span v-if="loading">Загрузка...</span>
  <template v-else>
    <p>
      <span @click="year -= 1">&lt;&lt; {{ year - 1 }}</span>
      {{ year }}
      <span @click="year += 1">{{ year + 1 }} &gt;&gt;</span>
    </p>

    <q-table
        v-if="rows.length > 0"
        :rows="rows"
        :columns="columns"
        row-key="a"
        hide-bottom
    />
    <p v-else>
      Нет платежей.
    </p>
  </template>

</template>
