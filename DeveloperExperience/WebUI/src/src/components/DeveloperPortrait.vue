<template>
  <div class="portrait">
    <el-row :gutter="20" class="firstrow">
      <el-col :span="7">
        <div class="grid-content bg-purple ele1">
          <div class="text-biaoti">贡献排名</div>
          <table class="table1">
            <tr class="first-tr">
              <td>
                <div class="data-progress">
                  <el-progress
                    type="dashboard"
                    :percentage="19"
                    :width="160"
                    :color="colors"
                  >
                    <template #default="{ percentage }">
                      <span class="percentage-value">{{ percentage }}%</span>
                      <span class="percentage-label">Issue</span>
                    </template>
                  </el-progress>
                </div>
              </td>

              <td>
                <div class="data-progress">
                  <el-progress
                    type="dashboard"
                    :percentage="57"
                    :width="160"
                    :color="colors"
                  >
                    <template #default="{ percentage }">
                      <span class="percentage-value">{{ percentage }}%</span>
                      <span class="percentage-label">PR/Commit</span>
                    </template>
                  </el-progress>
                </div>
              </td>
            </tr>
            <tr class="second-tr">
              <td colspan="2">
                <div
                  class="ShuJuTongJi"
                  id="ShuJuTongJi"
                  :style="{ width: '95%', height: '100%' }"
                ></div>
              </td>
            </tr>
          </table>
        </div>
      </el-col>
      <el-col :span="10">
        <div class="grid-content bg-purple ele2">
          <div class="text-biaoti">Labels</div>
          <img
            alt="logo"
            src="../assets/picture.jpg"
            style="width: 80%; height: 30%"
          />
          <div class="text-biaoti">协作网络</div>
          <div
            id="XieZuoGuanXi"
            :style="{ width: '100%', height: '60%' }"
          ></div>
        </div>
      </el-col>
      <el-col :span="7">
        <div class="grid-content bg-purple ele3">
          <div class="text-biaoti">
            开源影响力
            <div class="tipsright">
              <div class="tips">
                <a>2</a>
                <br />
                <span>Stars</span>
              </div>
              <div class="tips">
                <a>17</a>
                <br />
                <span>Watches</span>
              </div>
              <div class="tips">
                <a>0</a>
                <br />
                <span>Followers</span>
              </div>
              <div class="tips">
                <a>1</a>
                <br />
                <span>Following</span>
              </div>
            </div>
          </div>
          <div class="tableMoKuai">
            <table class="tabledata">
              <tr>
                <td colspan="2"><span>Ming's Gitee Stats</span></td>
              </tr>
              <tr>
                <td>Total Statrs Earned:</td>
                <td>34.5k</td>
              </tr>
              <tr>
                <td>Total Commits:</td>
                <td>8.2k</td>
              </tr>
              <tr>
                <td>Total PRs:</td>
                <td>491</td>
              </tr>
              <tr>
                <td>Total Issues:</td>
                <td>122</td>
              </tr>
              <tr>
                <td>Contributed to:</td>
                <td>14</td>
              </tr>
            </table>
          </div>

          <div class="tableMoKuai2">
            <div class="languagedata">
              <span>Most Used Language: Python</span>

              <el-progress
                :text-inside="true"
                :stroke-width="26"
                :color="colors"
                :percentage="85"
              />
            </div>
          </div>

          <div class="text-biaotisecond">Meetup/Meeting/Report</div>
          <button class="viewbutton">查看报告</button>
        </div>
      </el-col>
    </el-row>
    <el-row :gutter="20" class="secondrow">
      <el-col :span="17">
        <div class="grid-content bg-purple ele4">
          <div class="text-biaoti">活跃度概览</div>
          <div class="lookarea">
            <div
              id="HuoYueDuGaiLan"
              :style="{ width: '95%', height: '100%' }"
            ></div>
          </div>
        </div>
      </el-col>
      <el-col :span="7">
        <div class="grid-content bg-purple ele5">
          <div class="text-biaoti">NPS/NSS</div>
          <div id="NPSLeiDaTu" :style="{ width: '80%', height: '50%' }"></div>
          <table id="customers">
            <tr>
              <th>序号</th>
              <th>时间</th>
              <th>指标名称</th>
            </tr>
            <tr>
              <td>01</td>
              <td>20220101</td>
              <td>指标A</td>
            </tr>
            <tr class="alt">
              <td>02</td>
              <td>20220101</td>
              <td>指标B</td>
            </tr>
          </table>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import * as echarts from "echarts";
import { onMounted } from "vue";
import $ from "jquery";

let base = +new Date(2021, 12, 1);
let oneDay = 24 * 3600 * 1000;
let valueBase = Math.random() * 300;
let valueBase2 = Math.random() * 50;
let data = [];
let data2 = [];

for (var i = 1; i < 10; i++) {
  var now = new Date((base += oneDay));
  var dayStr = [now.getFullYear(), now.getMonth() + 1, now.getDate()].join("-");

  valueBase = Math.round((Math.random() - 0.5) * 20 + valueBase);
  valueBase <= 0 && (valueBase = Math.random() * 300);
  data.push([dayStr, valueBase]);

  valueBase2 = Math.round((Math.random() - 0.5) * 20 + valueBase2);
  valueBase2 <= 0 && (valueBase2 = Math.random() * 50);
  data2.push([dayStr, valueBase2]);
}
interface GraphNode {
  symbolSize: number;
  label?: {
    show?: boolean;
  };
}
export default defineComponent({
  name: "DeveloperPortrait",
  data() {
    return {
      percentage: 10,
      colors: [
        { color: "#FFB6C1", percentage: 20 },
        { color: "#ADADAD", percentage: 40 },
        { color: "#FFBB77", percentage: 60 },
        { color: "#97CBFF", percentage: 80 },
        { color: "#B8B8DC", percentage: 100 },
      ],
    };
  },
  setup() {
    onMounted(() => {
      let HuoYueDuGaiLan = echarts.init(
        document.getElementById("HuoYueDuGaiLan")
      );
      let NPSLeiDaTu = echarts.init(document.getElementById("NPSLeiDaTu"));
      let XieZuoGuanXi = echarts.init(document.getElementById("XieZuoGuanXi"));
      XieZuoGuanXi.showLoading();
      $.getJSON("/les-miserables.json", function (graph) {
        XieZuoGuanXi.hideLoading();

        graph.nodes.forEach(function (node: GraphNode) {
          node.label = {
            show: node.symbolSize > 25,
          };
        });
        XieZuoGuanXi.setOption({
          title: {
            top: "bottom",
            left: "right",
          },
          tooltip: {},
          legend: [
            {
              data: graph.categories.map(function (a: { name: string }) {
                return a.name;
              }),
              orient: "vertical",
              left: "10%",
              top: "15%",
            },
          ],
          animationDurationUpdate: 1500,
          animationEasingUpdate: "quinticInOut",
          series: [
            {
              name: "Les Miserables",
              type: "graph",
              radius: ["5%", "10%"],
              layout: "circular",
              circular: {
                rotateLabel: true,
              },
              data: graph.nodes,
              links: graph.links,
              categories: graph.categories,
              roam: true,
              label: {
                position: "right",
                formatter: "{b}",
              },
              lineStyle: {
                color: "source",
                curveness: 0.3,
              },
            },
          ],
        });
      });
      let ShuJuTongJi = echarts.init(document.getElementById("ShuJuTongJi"));

      HuoYueDuGaiLan.setOption({
        title: {
          left: "center",
        },
        legend: {
          top: "bottom",
          data: ["Intention"],
        },
        tooltip: {
          triggerOn: "none",
          position: function (pt) {
            return [pt[0], 130];
          },
        },
        toolbox: {
          left: "center",
          itemSize: 25,
          top: 55,
          feature: {
            dataZoom: {
              yAxisIndex: "none",
            },
            restore: {},
          },
        },
        xAxis: {
          type: "time",
          axisPointer: {
            value: "2016-10-7",
            snap: true,
            lineStyle: {
              color: "#7581BD",
              width: 2,
            },
            label: {
              show: true,
              formatter: function (params: any) {
                return echarts.format.formatTime("yyyy-MM-dd", params.value);
              },
              backgroundColor: "#7581BD",
            },
            handle: {
              show: true,
              color: "#7581BD",
            },
          },
          splitLine: {
            show: false,
          },
        },
        yAxis: {
          type: "value",
          axisTick: {
            inside: true,
          },
          splitLine: {
            show: false,
          },
          axisLabel: {
            inside: true,
            formatter: "{value}\n",
          },
          z: 10,
        },
        grid: {
          top: 110,
          left: 15,
          right: 15,
          height: 160,
        },
        dataZoom: [
          {
            type: "inside",
            throttle: 50,
          },
        ],
        series: [
          {
            name: "Fake Data",
            type: "line",
            smooth: true,
            symbol: "circle",
            symbolSize: 5,
            sampling: "average",
            itemStyle: {
              color: "#FDC147",
            },
            stack: "a",
            areaStyle: {
              color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                {
                  offset: 0,
                  color: "rgba(255,244,222,0.8)",
                },
                {
                  offset: 1,
                  color: "rgba(255,244,222,0.3)",
                },
              ]),
            },
            data: data,
          },
          {
            name: "Fake Data",
            type: "line",
            smooth: true,
            stack: "a",
            symbol: "circle",
            symbolSize: 5,
            sampling: "average",
            itemStyle: {
              color: "#78A6FC",
            },
            areaStyle: {
              color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                {
                  offset: 0,
                  color: "rgba(231,237,254,0.8)",
                },
                {
                  offset: 1,
                  color: "rgba(231,237,254,0.3)",
                },
              ]),
            },
            data: data2,
          },
        ],
      });
      const data1: number[] = [];
      for (let i = 0; i < 7; ++i) {
        data1.push(Math.round(Math.random() * 100));
      }

      NPSLeiDaTu.setOption({
        legend: {
          data: ["A", "B"],
          orient: "vertical",
          left: "80%",
          top: "40%",
        },
        radar: {
          indicator: [
            { name: "SIG", max: 6500 },
            { name: "发布", max: 16000 },
            { name: "本地开发", max: 30000 },
            { name: "PR", max: 38000 },
            { name: "ISSUE", max: 52000 },
          ],
        },
        series: [
          {
            type: "radar",
            radius: ["8%", "12%"],
            itemStyle: {
              normal: {
                color: function (params) {
                  var colorList = ["#898EFC", "#caf5f5"];
                  return colorList[params.dataIndex];
                },
              },
            },
            data: [
              {
                value: [4200, 3000, 20000, 35000, 50000],
                name: "A",
              },
              {
                value: [5000, 14000, 28000, 26000, 42000],
                name: "B",
              },
            ],
          },
        ],
      });
      ShuJuTongJi.setOption({
        xAxis: {
          show: false,
          type: "value",
          min: 0,
          max: 100,
        },
        yAxis: {
          axisLine: {
            show: false,
          },
          axisTick: {
            show: false,
          },
          type: "category",
          data: [
            "建立Issue",
            "尚未关闭Issue",
            "参与Issue讨论",
            "建立PR",
            "尚未关闭PR",
            "参与PR讨论",
            "合入commit",
          ],
          axisLabel: {
            inside: false,
            textStyle: {
              fontSize: "10",
              itemSize: "",
            },
          },
          inverse: true,
          animationDuration: 300,
          animationDurationUpdate: 300,
          max: 6,
        },
        grid: {
          left: "3%",
          right: "4%",
          bottom: "3%",
          containLabel: true,
          show: false,
        },
        series: [
          {
            realtimeSort: true,
            name: "贡献排名",
            type: "bar",
            itemStyle: {
              normal: {
                color: function (params) {
                  var colorList = [
                    "#898EFC",
                    "#898EFC",
                    "#898EFC",
                    "#898EFC",
                    "#898EFC",
                    "#898EFC",
                    "#898EFC",
                    "#898EFC",
                  ];
                  return colorList[params.dataIndex];
                },
              },
            },
            barWidth: "25%",
            radius: ["5%", "10%"],
            data: data1,
            label: {
              show: true,
              position: "right",
              valueAnimation: true,
            },
          },
        ],
        legend: {
          show: false,
          bottom: 0,
        },
        animationDuration: 0,
        animationDurationUpdate: 3000,
        animationEasing: "linear",
        animationEasingUpdate: "linear",
      });
      window.onresize = function () {
        HuoYueDuGaiLan.resize();
        XieZuoGuanXi.resize();
        ShuJuTongJi.resize();
        NPSLeiDaTu.resize();
      };
    });
  },
});
</script>

<style scoped>
.portrait {
  position: absolute;
  top: 80px;
  left: 230px;
  bottom: 0px;
  right: 0px;
  padding: 10px;
  overflow-y: auto;
  background-color: rgb(245, 245, 245);
}

h3 {
  margin: 40px 0 0;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}

.pagemain {
  position: absolute;
  top: 50px;
  left: 230px;
  bottom: 0px;
  right: 0px;
  padding: 10px;
  overflow-y: auto;
  background-color: rgb(245, 245, 245);
}

.el-row:last-child {
  margin-bottom: 0;
}

.el-col {
  border-radius: 4px;
}

.bg-purple-dark {
  background: #fcfcfd;
}

.bg-purple {
  background: #fff;
}

.bg-purple-light {
  background: #e5e9f2;
}

.grid-content {
  border-radius: 4px;
  min-height: 36px;
}

.firstrow {
  height: 61%;
}

.secondrow {
  height: 30%;
}

.ele1,
.ele2,
.ele3 {
  height: 90%;
}

.ele4,
.ele5 {
  height: 100%;
}

.table1 {
  width: 100%;
  height: 100%;
  padding-top: 10px;
  padding-bottom: 0px;
}

.el-progress {
  padding-top: 12px;
}

.first-tr {
  height: 5%;
}

.ShuJuTongJi {
  margin-top: -10%;
  margin-left: 0%;
}

.text-biaoti {
  padding-top: 10px;
  padding-left: 10px;
  text-align: left;
  color: rgb(156, 155, 155);
}

.text-biaotisecond {
  font-size: 20px;
  padding-top: 10px;
  padding-left: 10px;
  padding-bottom: 20px;
  text-align: left;
  color: rgb(156, 155, 155);
}

.tips {
  display: inline-block;
  padding-right: 20px;
  text-align: center;
}

.tips a {
  font-size: 25px;
  color: #23718e;
}

.tipsright {
  float: right;
}

.tableMoKuai {
  margin-top: 40px;
  text-align: left;
  margin-left: 10px;
  margin-bottom: 10px;
  border: 1px solid rgb(211, 210, 210);
  border-radius: 4px;
  padding-top: 10px;
  padding-left: 10px;
  padding-bottom: 10px;
}

.tableMoKuai2 {
  margin-top: 5px;
  text-align: left;
  margin-left: 10px;
  margin-bottom: 15px;
  border: 1px solid rgb(211, 210, 210);
  border-radius: 4px;
  padding-top: 10px;
  padding-left: 10px;
  padding-bottom: 10px;
}

.tabledata td {
  width: 150px;
}

.tabledata tr {
  height: 30px;
}

.tabledata {
  padding-left: 10px;
}

.languagedata {
  padding-left: 10px;
  padding-right: 10px;
}

.tableMoKuai span {
  font-size: 22px;
  color: rgb(82, 151, 229);
  margin-bottom: 20px;
}

.tableMoKuai2 span {
  color: rgb(82, 151, 229);
  font-size: 22px;
  font-weight: 200;
}

.viewbutton {
  background-color: #878ffe;
  border: none;
  color: white;
  padding: 13px 95px;
  text-align: center;
  border-radius: 4px;
  text-decoration: none;
  display: inline-block;
  font-size: 18px;
  margin: 4px 2px;
  cursor: pointer;
}

.percentage-value {
  display: block;
  margin-top: 10px;
  font-size: 28px;
}

.percentage-label {
  display: block;
  margin-top: 10px;
  font-size: 12px;
}

#customers {
  width: 90%;
  border-collapse: collapse;
  margin-left: 5%;
}

#customers td,
#customers th {
  font-size: 14px;
  text-align: center;
  border: 1px solid #dddddc;
  border-left: none;
  border-right: none;
  border-top: none;
  padding: 10px 7px 10px 7px;
}

#customers th {
  font-size: 16px;
  text-align: center;
  padding-top: 5px;
  padding-bottom: 4px;
}

#customers tr.alt td {
  color: #000000;
  background-color: #fdfdfb;
}

.lookarea {
  display: inline;
}
</style>
