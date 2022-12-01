(require '[clojure.string :as str])

(def read-input
  (slurp (first *command-line-args*)))

(defn parse-int
  [x]
  (Integer/parseInt x))

(defn elf-calories
  [elf]
  (->> elf
       (str/split-lines)
       (map parse-int)
       (reduce +)))

(defn part-1
  [input]
  (->> (str/split input #"\n\n")
       (map elf-calories)
       (apply max)))

(defn part-2
  [input]
  (->> (str/split input #"\n\n")
       (map elf-calories)
       (sort >)
       (take 3)
       (reduce +)))

(def input read-input)
(println "Part 1:" (part-1 input))
(println "Part 2:" (part-2 input))
