import sys
import os
import pandas as pd
import seaborn as sns
import matplotlib.pyplot as plt

def main():
    if len(sys.argv) < 2:
        print("❌ Укажите путь к CSV-файлу")
        sys.exit(1)

    csv_path = sys.argv[1]
    if not os.path.exists(csv_path):
        print(f"❌ Файл не найден: {csv_path}")
        sys.exit(1)

    date_part = os.path.splitext(os.path.basename(csv_path))[0]
    output_dir = os.path.join("charts")
    os.makedirs(output_dir, exist_ok=True)

    output_path = os.path.join(output_dir, f"{date_part}.png")
    if os.path.exists(output_path):
        print(f"✅ Визуализация уже существует: {output_path}")
        return

    df = pd.read_csv(csv_path)
    df["Курс"] = df["Курс"].astype(float).round(2)

    plt.figure(figsize=(14, 9))
    sns.set(style="whitegrid")

    bars = sns.barplot(
        x="Курс",
        y="Валюта",
        data=df,
        color="skyblue"
    )

    date_part = os.path.splitext(os.path.basename(csv_path))[0]
    title = f"Курсы валют к тенге на {date_part}"

    plt.title(title, fontsize=18)
    plt.xlabel("Курс (в тенге)")
    plt.ylabel("Валюта")

    for bar in bars.containers[0]:
        width = bar.get_width()
        bars.text(
            width + 1,
            bar.get_y() + bar.get_height() / 2,
            f"{width:.2f}",
            va="center",
            ha="left",
            fontsize=10,
            color="black"
        )

    plt.tight_layout()
    plt.savefig(output_path)
    print(f"📊 График сохранён: {output_path}")

if __name__ == "__main__":
    main()
