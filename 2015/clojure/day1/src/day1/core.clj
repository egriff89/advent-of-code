(ns day1.core)

(def steps (slurp "resources/input.txt"))

(defn check-directions [current-step instruction]
  (= 0 (compare (str current-step) instruction)))

(defn which-floor? [input]
  (let [floor 0]
    (reduce + (map (fn [step]
           (if (check-directions step "(")
             (inc floor)
             (dec floor))) input))))

(defn first-basement-visit? [input]
  (loop [floor 0 position 0 steps input]
    (cond
      (= floor -1) position
      (check-directions (first steps) "(") (recur (inc floor) (inc position) (rest steps))
      (check-directions (first steps) ")") (recur (dec floor) (inc position) (rest steps)))))


(defn -main []
  ;; Part 1
  (println (which-floor? steps))
  ;; Part 2
  (println (first-basement-visit? steps)))