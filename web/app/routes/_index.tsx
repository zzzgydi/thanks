import { useNavigate } from "@remix-run/react";
import { useLockFn } from "ahooks";
import { useState } from "react";
import { Button } from "~/components/ui/button";
import { Input } from "~/components/ui/input";
import { LangSettings } from "~/services/constants";
import confetti from "canvas-confetti";
import GithubSvg from "~/assets/github.svg?react";

export default function Index() {
  const [lang, setLang] = useState("node");
  const [url, setUrl] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const navigate = useNavigate();

  const handleSubmit = useLockFn(async () => {
    if (!url || loading) return;
    setLoading(true);

    confetti({ particleCount: 100, spread: 70, origin: { y: 0.6 } });

    try {
      const result = await fetch("/api/task/create", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ lang, url }),
      }).then((res) => res.json());

      if (result.code !== 0) {
        throw new Error(result.msg);
      }
      navigate(`/t/${result.data.id}`);
    } catch (error: any) {
      console.error(error);
      setError(error.message);
    } finally {
      setLoading(false);
    }
  });

  const currentLang = LangSettings.find((i) => i.lang === lang);

  return (
    <div className="w-screen min-h-dvh noise-bg flex flex-col items-center justify-center">
      <h1 className="main-text px-6 mt-[10vh]">OpenSource Thanks</h1>

      <div className="w-full max-w-[750px] px-8 py-8">
        <div className="space-y-2 pb-8">
          <div className="flex items-center gap-2">
            {LangSettings.map((i) => (
              <div
                key={i.lang}
                className="flex items-center gap-1"
                onClick={() => setLang(i.lang)}
              >
                <input
                  type="radio"
                  name="lang"
                  value={i.lang}
                  checked={lang === i.lang}
                  onChange={() => setLang(i.lang)}
                />
                <label className="cursor-pointer">{i.label}</label>
              </div>
            ))}
          </div>

          <div className="flex items-center gap-2">
            <Input
              value={url}
              onChange={(e) => setUrl(e.target.value)}
              placeholder={currentLang?.placeholder}
            />

            <Button onClick={handleSubmit} disabled={loading}>
              {loading ? "Creating" : "Create"}
            </Button>
          </div>

          <div className="text-sm text-muted-foreground space-y-1 overflow-auto">
            <p>
              Provide the link to your project's{" "}
              <span className="font-medium">
                {currentLang?.file ?? "package.json"}
              </span>{" "}
              to get a list of contributors and their contributions.
            </p>
            <p>
              e.g.,{" "}
              <span className="font-medium">
                {currentLang?.example ?? "https://path/to/your/package.json"}
              </span>
            </p>
            <p>Note: Contribution amounts are for reference only.</p>
          </div>
        </div>

        {error && (
          <div className="w-full max-h-[50vh] overflow-auto">
            <pre className="text-destructive">Error: {error}</pre>
          </div>
        )}
      </div>

      <footer className="mt-auto flex items-center py-2 gap-2 px-6">
        <div className="text-sm text-muted-foreground">
          Copyright © 2024 OpenSourceThanks. All rights reserved.
        </div>

        <a href="https://github.com/zzzgydi/thanks">
          <GithubSvg className="w-6 h-6" />
        </a>
      </footer>
    </div>
  );
}
