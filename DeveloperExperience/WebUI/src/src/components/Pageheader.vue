<template>
  <div class="header">
    <el-row>
      <div class="header-icon">
        <el-button type="primary" plain>文件</el-button>
        <el-button type="success" plain>模型</el-button>
        <el-button type="info" plain>分析</el-button>
        <el-button type="warning" plain>管理</el-button>
        <el-button type="danger" plain>成员</el-button>
        <el-button type="info" plain>添加</el-button>
      </div>
    </el-row>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "Pageheader",
});
</script>

<style>
.header {
  position: absolute;
  top: 0px;
  left: 230px;
  right: 0px;
  background-color: #fff;
  padding-top: 5px;
}
.header-icon {
  margin-left: 5px;
  margin-top: 10px;
}
.el-button {
  height: 40px;
}
.rightDropdown {
  position: fixed;
  right: 15px;
  top: 18px;
}
.rightDropdown span {
  color: rgb(156, 155, 155);
  font-size: 15px;
}

.admin-text {
  margin-bottom: 10px;
}
</style>