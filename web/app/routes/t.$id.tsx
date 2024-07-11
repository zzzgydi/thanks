import { LoaderFunctionArgs } from "@remix-run/node";
import { Link, useLoaderData, useParams } from "@remix-run/react";
import { useLockFn } from "ahooks";
import { useState } from "react";
import { Virtuoso } from "react-virtuoso";
import { IThkTaskResponse } from "~/services/types";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "~/components/ui/tooltip";
import GithubSvg from "~/assets/github.svg?react";

export const loader = async ({ params, request }: LoaderFunctionArgs) => {
  const { id } = params;
  const url = new URL(request.url);

  const result = await fetch(`${url.origin}/api/task/${id}`).then((res) =>
    res.json()
  );

  return { task: result?.data as IThkTaskResponse };
};

export default function TaskDetail() {
  const { id } = useParams();
  const { task } = useLoaderData<ReturnType<typeof loader>>();

  const [list, setList] = useState(task.list);

  const loadMore = useLockFn(async () => {
    if (list.length >= task.total) return;
    const result = await fetch(`/api/task/${id}?offset=${list.length}`).then(
      (res) => res.json()
    );
    setList((prev) => [...prev, ...result.data.list]);
  });

  return (
    <div className="w-full h-screen h-dvh noise-bg overflow-hidden relative">
      <div className="absolute inset-x-0 backdrop-blur-md z-20">
        <div className="w-full max-w-[750px] px-6 py-2 mx-auto flex-none flex items-center">
          <h1 className="border-text text-3xl">
            <Link to="/" className="border-text">
              OpenSource Thanks
            </Link>
          </h1>

          <a href="https://github.com/zzzgydi/thanks" className="ml-auto mr-2">
            <GithubSvg className="w-5 h-5" />
          </a>
        </div>
      </div>

      <div className="h-full">
        <Virtuoso
          components={{
            Header: () => <div className="w-full h-[60px]" />,
            Footer: () =>
              list.length >= task.total ? (
                <div className="py-5 text-center text-muted-foreground">
                  End of List
                </div>
              ) : null,
          }}
          endReached={list.length >= task.total ? undefined : loadMore}
          totalCount={list.length}
          itemContent={(index) => {
            const c = list[index];
            return (
              <div
                key={c.id}
                className="max-w-[750px] mx-auto px-6 flex items-center gap-2 mb-1.5"
              >
                <img
                  className="w-8 h-8 rounded-full"
                  src={`https://avatars.githubusercontent.com/u/${c.id}?v=4`}
                  alt=""
                />
                <div className="text-md mr-auto">
                  <a
                    href={`https://github.com/${c.login}`}
                    className="hover:underline font-medium text-primary"
                  >
                    {c.login}
                  </a>
                </div>

                <div>
                  <TooltipProvider>
                    <Tooltip>
                      <TooltipTrigger>
                        <div className="text-sm min-w-6 h-6 bg-muted/50 rounded flex items-center justify-center">
                          {c.repos.length}
                        </div>
                      </TooltipTrigger>
                      <TooltipContent>
                        {c.repos.map((r) => (
                          <div key={r.repo} className="flex items-center gap-4">
                            <span className="font-medium">{r.repo}</span>
                            <span className="ml-auto">
                              {(r.score * 1000).toFixed(2)}&permil;
                            </span>
                          </div>
                        ))}
                      </TooltipContent>
                    </Tooltip>
                  </TooltipProvider>
                </div>

                <div>
                  <div
                    className="w-[80px] md:w-[120px] h-6 rounded relative bg-muted/50 overflow-hidden"
                    title={`Contributions: ${c.total * 1000}‰`}
                  >
                    <div
                      className="h-full absolute right-0 top-0 bg-orange-400/90"
                      style={{ width: `${Math.min(100, c.total * 1000)}%` }}
                    />

                    <div className="absolute text-sm text-primary inset-0 flex items-center justify-center">
                      {(c.total * 1000).toFixed(2)}&permil;
                    </div>
                  </div>
                </div>
              </div>
            );
          }}
        />
      </div>
    </div>
  );
}
