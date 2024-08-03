<script setup lang="ts">
import { useRouter } from 'vue-router';
import { h, onMounted, ref, Text } from "vue";
import type { DataTableColumns } from "naive-ui";
import { useMessage,NButton } from "naive-ui";
import { TaskAccount, TaskAccountItem } from "@/api/task/task";
import { Ok } from "@/api/task/common";

const router = useRouter();

const data = ref<TaskAccountItem[]>([]);

const message = useMessage();
const columns = createColumns({
  play(row: TaskAccountItem) {
    message.info(`Play ${row.id}`);

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
</template>

<style scoped lang="less">

</style>
