import { NavLink } from "react-router-dom";
import Table from "../../components/Table";

export default function IncidentTable({ data }) {
  const columns = [
    { accessorKey: "id", header: "ID" },
    {
      accessorKey: "title",
      header: "Title",
      cell: (info) => (
        <NavLink
          to={`/incidents/${info.row.original.id}`}
          style={{ color: "#1976d2", textDecoration: "none", fontWeight: 500 }}
        >
          {info.getValue()}
        </NavLink>
      ),
    },
    { accessorKey: "ai_severity", header: "Severity" },
    { accessorKey: "ai_category", header: "Category" },
    { accessorKey: "affected_service", header: "Service" },
    { accessorKey: "created_at", header: "Created At" },
  ];

  return <Table columns={columns} data={data} />;
}
