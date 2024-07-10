import { useParams } from "@remix-run/react";

export default function TaskDetail() {
  const { id } = useParams();

  return (
    <div>
      <h1>Task Detail</h1>
      <p>Task ID: {id}</p>
    </div>
  );
}
