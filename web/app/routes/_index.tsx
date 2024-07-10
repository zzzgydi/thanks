import type { MetaFunction } from "@remix-run/node";
import { Link } from "@remix-run/react";
import { useLockFn } from "ahooks";
import { useState } from "react";
import { Button } from "~/components/ui/button";
import { Input } from "~/components/ui/input";

export const meta: MetaFunction = () => {
  return [
    { title: "OpenSource Thanks" },
    { name: "description", content: "Thank You for Your Contributions" },
  ];
};

export default function Index() {
  const [lang, setLang] = useState("node");
  const [minScore, setMinScore] = useState(0.0001);
  const [url, setUrl] = useState(
    "https://github.com/zzzgydi/thanks/raw/main/web/package.json"
  );
  const [loading, setLoading] = useState(false);
  const [resultId, setResultId] = useState<string | null>(null);

  const handleSubmit = useLockFn(async () => {
    if (!url || loading) return;
    setLoading(true);

    try {
      const result = await fetch("/api/task/create", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          lang,
          url,
          min_score: Math.max(minScore, 0.0001),
        }),
      }).then((res) => res.json());
      console.log(result);
      setResultId(result.data.id);
    } catch (error: any) {
      console.error(error);
    } finally {
      setLoading(false);
    }
  });

  return (
    <div className="w-screen min-h-dvh noise-bg flex flex-col items-center justify-center">
      <h1 className="main-text px-6">OpenSource Thanks</h1>

      <div className="w-full max-w-[750px] px-8 py-8 mb-[100px] space-y-2">
        <div className="flex items-center gap-2">
          {["node", "golang"].map((i) => (
            <div
              key={i}
              className="flex items-center gap-1"
              onClick={() => setLang(i)}
            >
              <input
                type="radio"
                name="lang"
                value={i}
                checked={lang === i}
                onChange={() => setLang(i)}
              />
              <label className="cursor-pointer">{i}</label>
            </div>
          ))}
        </div>

        <div className="flex items-center gap-2">
          <span>Min Score</span>
          <Input
            className="max-w-[100px]"
            type="number"
            value={minScore}
            onChange={(e) => setMinScore(Number(e.target.value))}
          />
        </div>

        <div className="flex items-center gap-2">
          <Input
            value={url}
            onChange={(e) => setUrl(e.target.value)}
            placeholder={
              lang === "node"
                ? "https://path/to/your/package.json"
                : "https://path/to/your/go.mod"
            }
          />

          <Button onClick={handleSubmit} disabled={loading}>
            {loading ? "Creating" : "Create"}
          </Button>
        </div>
        {resultId && (
          <div>
            <p className="mt-4">Your Task ID is: {resultId}</p>
            <p className="mt-4">
              You can check the progress{" "}
              <Link to={`/t/${resultId}`} className="text-blue-500">
                /t/{resultId}
              </Link>
            </p>
          </div>
        )}
      </div>
    </div>
  );
}
