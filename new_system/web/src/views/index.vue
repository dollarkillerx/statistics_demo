<script setup lang="ts">
import { useRouter } from 'vue-router';
import { h, onMounted, ref, Text } from "vue";
import type { DataTableColumns } from "naive-ui";
import { NButton } from "naive-ui";
import { GetTaskByAccount, TaskAccount,  } from "@/api/task";
import type {Task,TaskAccountItem} from "@/api/task";
import { Ok } from "@/api/common";

const data = ref<TaskAccountItem[]>([]);
const tasks = ref<Task>()
const columns = createColumns({
  async play(row: TaskAccountItem) {
    // message.info(`Play ${row.id}`);
    const resp = await GetTaskByAccount(row.client_id);
    if (resp instanceof Ok) {
      tasks.value = resp.value;
      console.log(tasks.value);
    }
  }
});

function createColumns({
                         play
                       }: {
  play: (row: TaskAccountItem) => void
}): DataTableColumns<TaskAccountItem> {
  return [
    {
      title: "Id",
      key: "client_id"
    },
    {
      title: "Account",
      key: "account"
    },
    {
      title: "杠杆",
      key: "leverage"
    },
    {
      title: "服务器",
      key: "server"
    },
    {
      title: "经纪商",
      key: "company"
    },
    {
      title: "余额",
      key: "balance",
      render(row) {
        return "$" + row.balance.toFixed(2)
      }
    },
    {
      title: "持仓盈亏",
      key: "profit",
      render(row) {
        return "$" +row.profit.toFixed(2)
      }
    },
    {
      title: "统计分析",
      key: "id",
      render(row) {
        return h(
          NButton,
          {
            strong: true,
            tertiary: true,
            size: "small",
            onClick: () => play(row)
          },
          { default: () => "统计分析" }
        );
      }
    }
  ];
}

onMounted(async () => {
  const resp = await TaskAccount();
  if (resp instanceof Ok) {
    data.value = resp.value;
    console.log(data.value);
  }
});

const formatDate = (timestamp: number) => {
  const date = new Date(timestamp * 1000);
  const year = date.getFullYear();
  const month = ('0' + (date.getMonth() + 1)).slice(-2);
  const day = ('0' + date.getDate()).slice(-2);
  const hours = ('0' + date.getHours()).slice(-2);
  const minutes = ('0' + date.getMinutes()).slice(-2);
  const seconds = ('0' + date.getSeconds()).slice(-2);
  return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
}

</script>

<template>
  <div>
    <n-data-table
      :columns="columns"
      :data="data"
      :pagination="false"
      :bordered="false"
    />
  </div>

  <div v-if="tasks" class="bg-white my-3 rounded-xl p-1">
    <div class="text-2xl ">数据统计: </div>
    <div v-for="item in tasks.profits" :key="item.period">
      <div v-if="item.min_profit != 0 && item.max_profit != 0" class="flex flex-row space-x-10">
        <div> 时间： {{ item.period }}</div>
        <div>最低利润: {{ item.min_profit }}</div>
        <div>最高利润: {{ item.max_profit }}</div>
      </div>
    </div>


    <div v-if="tasks" class="bg-white my-3 rounded-xl p-1">
      <div class="text-2xl ">当前持仓: </div>
      <n-table :bordered="false" :single-line="false">
        <thead>
        <tr>
          <th>order id</th>
          <th>opening_time</th>
          <th>direction</th>
          <th>symbol</th>
          <th>open price</th>
          <th>volume</th>
          <th>market</th>
          <th>profit</th>
          <th>common</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="item in tasks.positions" :key="item.id">
          <td>{{ item.order_id }}</td>
          <td>{{ formatDate(item.opening_time) }}</td>
          <td>{{ item.direction }}</td>
          <td>{{ item.symbol }}</td>
          <td>{{ item.open_price }}</td>
          <td>{{ item.volume }}</td>
          <td>{{ item.market }}</td>
          <td>{{ item.profit }}</td>
          <td>{{ item.common }}</td>
        </tr>

        </tbody>
      </n-table>
    </div>
  </div>
</template>

<style scoped lang="less">

</style>
