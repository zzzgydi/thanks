import { LoaderFunctionArgs } from "@remix-run/node";
import { useLoaderData, useParams } from "@remix-run/react";
import { IThkContributor } from "~/services/types";

export const loader = async ({ params, request }: LoaderFunctionArgs) => {
  const { id } = params;
  const url = new URL(request.url);

  const result = await fetch(`${url.origin}/api/task/detail`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ id }),
  }).then((res) => res.json());

  return { task: result?.data as { contributions: IThkContributor[] } };
};

export default function TaskDetail() {
  const { id } = useParams();
  const { task } = useLoaderData<ReturnType<typeof loader>>();

  return (
    <div className="w-full min-h-dvh noise-bg overflow-x-hidden">
      <p>Task ID: {id}</p>

      <div className="max-w-[750px] mx-auto px-6">
        {task?.contributions.map((contributor) => (
          <div key={contributor.id} className="flex items-center gap-2 mb-1">
            <img
              className="w-8 h-8 rounded-full"
              src={`https://avatars.githubusercontent.com/u/${contributor.id}?v=4`}
              alt=""
            />
            <div>{contributor.login}</div>
            <div>{contributor.total}</div>
          </div>
        ))}
      </div>
    </div>
  );
}