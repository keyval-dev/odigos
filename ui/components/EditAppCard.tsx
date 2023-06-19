import { ApplicationData } from "@/types/apps";
import { getLangIcon } from "@/utils/icons";
import Link from "next/link";
import { useState } from "react";

export default function EditAppCard({
  id,
  name,
  namespace,
  languages,
  kind,
  instrumented,
}: ApplicationData) {
  return (
    <div className="shadow-lg border border-gray-200 rounded-lg bg-white hover:bg-gray-100 cursor-pointer">
      <Link
        href={`/source/${namespace}/${kind.toString().toLowerCase()}/${name}`}
        className="flex flex-row p-3 items-center space-x-4">

        {getLangIcon(languages[0], "w-12 h-12", kind.toString())}
        <div className="flex flex-col items-start">
          <div className="font-bold">{name}</div>
          <div>{kind}</div>
          <div>namespace: {namespace}</div>
          {instrumented ? (
            <div className="text-green-600 font-bold">Instrumented</div>
          ) : (
            <div className="text-orange-400 font-bold">Not Instrumented</div>
          )}
        </div>

      </Link>
    </div>
  );
}
