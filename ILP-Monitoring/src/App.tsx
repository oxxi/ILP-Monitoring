import {
  DataGrid,
  HeaderFilter,
  Column,
  Sorting,
  Paging,
  GroupPanel,
  Export,
} from 'devextreme-react/data-grid';
import 'devextreme/dist/css/dx.light.css';
import { exportDataGrid } from 'devextreme/excel_exporter';
import { useGetData } from './hook/useData';
import { useEffect } from 'react';
import { Workbook } from 'exceljs';
import { saveAs } from 'file-saver-es';

function App() {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const onHandleExporting = (data: any) => {
    const workbook = new Workbook();
    const worksheet = workbook.addWorksheet('Main sheet');

    exportDataGrid({
      component: data.component,
      worksheet: worksheet,
      autoFilterEnabled: true,
    }).then(() => {
      workbook.xlsx.writeBuffer().then((buffer) => {
        saveAs(
          new Blob([buffer], { type: 'application/octet-stream' }),
          'monitoreo_ilp.xlsx'
        );
      });
    });
    data.cancel = true;
  };

  const { data, loadData } = useGetData();

  useEffect(() => {
    loadData();
  }, []);

  const sortDataDescending = (dataArray: GetData[]): GetData[] => {
    if (data.length === 0) return [];

    return dataArray.sort((a, b) => {
      if (!a.dateInit || !b.dateInit) {
        return 0;
      }
      return b.dateInit.getTime() - a.dateInit.getTime();
    });
  };

  return (
    <>
      <section className="w-100 p-4">
        <h1 className="text-base text-gray-600 font-semibold p-2">
          Monitoreo de Recursos APP ILP
        </h1>

        <DataGrid
          width="100%"
          allowColumnReordering={true}
          allowColumnResizing={true}
          columnHidingEnabled={true}
          onExporting={onHandleExporting}
          rowAlternationEnabled={true}
          dataSource={sortDataDescending(data)}
          showBorders={true}
        >
          <HeaderFilter visible={true} />
          <GroupPanel visible={true} />
          <Sorting mode="multiple" />
          <Column dataField="resource" caption="Recurso" alignment="left" />
          <Column
            dataField="status"
            width={150}
            caption="Estado"
            alignment="center"
          />
          <Column
            dataField="statusCode"
            width={150}
            caption="Código estado"
            alignment="center"
          />
          <Column
            name="date_filter"
            width={150}
            sortOrder="desc"
            dataField="timeStart"
            dataType="date"
            format={'dd/MM/yyyy'}
            caption="Fecha"
            alignment="center"
          />
          <Column
            width={150}
            dataField="timeStart"
            dataType="date"
            caption="Inicio"
            format={'hh:mm:ss a'}
            alignment="center"
          />
          <Column
            width={150}
            dataField="timeEnd"
            dataType="datetime"
            format={'hh:mm:ss a'}
            caption="Fin"
            alignment="center"
          />
          <Column
            width={150}
            dataField="duration"
            caption="Duración (Milisegundos)"
            alignment="center"
          />
          <Paging defaultPageSize={20} />
          <Export enabled={true} />
        </DataGrid>
      </section>
    </>
  );
}

export default App;
